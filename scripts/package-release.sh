#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
DIST_DIR="${1:-$ROOT_DIR/dist}"
DEFAULT_VERSION="$(sed -n 's/^var Value = "\(.*\)"$/\1/p' "$ROOT_DIR/internal/version/version.go" | head -1)"
VERSION="${VERSION:-${GITHUB_REF_NAME:-${DEFAULT_VERSION:-dev}}}"

mkdir -p "$DIST_DIR"
rm -f "$DIST_DIR"/zaibflow_*.tar.gz "$DIST_DIR"/zaibflow_*.zip "$DIST_DIR"/checksums.txt "$DIST_DIR"/latest.json

build_tar_target() {
  local os="$1"
  local arch="$2"
  local work="$DIST_DIR/${os}-${arch}"
  local asset="zaibflow_${os}_${arch}.tar.gz"

  rm -rf "$work"
  mkdir -p "$work"
  GOOS="$os" GOARCH="$arch" CGO_ENABLED=0 go build \
    -trimpath \
    -ldflags="-s -w -X github.com/Shahzaibzah00r/zaibflow/internal/version.Value=${VERSION}" \
    -o "$work/zaibflow" \
    ./cmd/zaibflow
  tar -C "$work" -czf "$DIST_DIR/$asset" zaibflow
}

build_zip_target() {
  local os="$1"
  local arch="$2"
  local work="$DIST_DIR/${os}-${arch}"
  local asset="zaibflow_${os}_${arch}.zip"

  rm -rf "$work"
  mkdir -p "$work"
  GOOS="$os" GOARCH="$arch" CGO_ENABLED=0 go build \
    -trimpath \
    -ldflags="-s -w -X github.com/Shahzaibzah00r/zaibflow/internal/version.Value=${VERSION}" \
    -o "$work/zaibflow.exe" \
    ./cmd/zaibflow
  (cd "$work" && ARCHIVE_PATH="$DIST_DIR/$asset" python3 - <<'PY'
import os
import pathlib
import zipfile

archive = pathlib.Path(os.environ["ARCHIVE_PATH"])
with zipfile.ZipFile(archive, "w", compression=zipfile.ZIP_DEFLATED) as zf:
    zf.write("zaibflow.exe", arcname="zaibflow.exe")
PY
  )
}

build_tar_target darwin amd64
build_tar_target darwin arm64
build_tar_target linux amd64
build_tar_target linux arm64
build_zip_target windows amd64
build_zip_target windows arm64

TAG_VERSION="$VERSION"
[[ "$TAG_VERSION" != v* ]] && TAG_VERSION="v$TAG_VERSION"

if command -v shasum >/dev/null 2>&1; then
  (cd "$DIST_DIR" && shasum -a 256 zaibflow_*.tar.gz zaibflow_*.zip > checksums.txt)
else
  (cd "$DIST_DIR" && sha256sum zaibflow_*.tar.gz zaibflow_*.zip > checksums.txt)
fi

cat > "$DIST_DIR/latest.json" <<EOF
{
  "version": "$TAG_VERSION",
  "url": "https://github.com/Shahzaibzah00r/zaibflow/releases/tag/$TAG_VERSION"
}
EOF
