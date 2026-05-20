#!/usr/bin/env bash
set -euo pipefail

REPO="Shahzaibzah00r/zaibflow"
VERSION="${ZAIBFLOW_VERSION:-latest}"
RELEASE_BASE_URL="${ZAIBFLOW_RELEASE_BASE_URL:-}"
TMP_DIR="$(mktemp -d)"
trap 'rm -rf "$TMP_DIR"' EXIT

detect_os() {
  case "$(uname -s)" in
    Darwin) echo "darwin" ;;
    Linux) echo "linux" ;;
    *) echo "unsupported operating system: $(uname -s)" >&2; exit 1 ;;
  esac
}

detect_arch() {
  case "$(uname -m)" in
    x86_64|amd64) echo "amd64" ;;
    arm64|aarch64) echo "arm64" ;;
    *) echo "unsupported architecture: $(uname -m)" >&2; exit 1 ;;
  esac
}

download_url() {
  local asset="$1"

  if [[ -n "$RELEASE_BASE_URL" ]]; then
    echo "${RELEASE_BASE_URL%/}/${asset}"
    return
  fi

  if [[ "$VERSION" == "latest" ]]; then
    echo "https://github.com/${REPO}/releases/latest/download/${asset}"
  else
    echo "https://github.com/${REPO}/releases/download/${VERSION}/${asset}"
  fi
}

verify_checksum() {
  local asset_path="$1"
  local checksums_path="$2"
  local asset_name
  local checksum_line
  local checksum_hash

  asset_name="$(basename "$asset_path")"

  if [[ ! -f "$checksums_path" ]]; then
    echo "warning: checksums.txt not found, skipping verification" >&2
    return 0
  fi

  checksum_line="$(grep -E "(^|[[:space:]/])${asset_name}$" "$checksums_path" || true)"

  if [[ -z "$checksum_line" ]]; then
    echo "warning: checksum for ${asset_name} not found, skipping verification" >&2
    return 0
  fi

  checksum_hash="$(echo "$checksum_line" | awk '{print $1}')"

  if [[ -z "$checksum_hash" ]]; then
    echo "warning: checksum hash for ${asset_name} is empty, skipping verification" >&2
    return 0
  fi

  if command -v shasum >/dev/null 2>&1; then
    echo "${checksum_hash}  ${asset_name}" | (cd "$(dirname "$asset_path")" && shasum -a 256 -c -)
  elif command -v sha256sum >/dev/null 2>&1; then
    echo "${checksum_hash}  ${asset_name}" | (cd "$(dirname "$asset_path")" && sha256sum -c -)
  else
    echo "warning: no checksum tool found, skipping verification" >&2
  fi
}

ensure_claude() {
  if command -v claude >/dev/null 2>&1; then
    return 0
  fi

  echo "Claude Code CLI is required."
  echo ""

  local os_name
  os_name="$(detect_os)"

  if [[ "$os_name" == "darwin" || "$os_name" == "linux" ]]; then
    echo "Installing Claude Code CLI automatically..."
    curl -fsSL https://claude.ai/install.sh | bash
    if command -v claude >/dev/null 2>&1; then
      return 0
    fi
    echo "error: Claude Code CLI installation failed or is not on PATH." >&2
    exit 1
  fi

  echo "Please install Claude Code CLI manually:"
  echo "  macOS/Linux: curl -fsSL https://claude.ai/install.sh | bash"
  echo "  Windows:     See https://docs.anthropic.com/en/docs/agents-and-tools/claude-code/overview"
  exit 1
}

get_install_dir() {
  local home="$1"
  if [[ -d "$home/bin" ]]; then
    echo "$home/bin"
    return
  fi
  if mkdir -p "$home/bin" 2>/dev/null; then
    echo "$home/bin"
    return
  fi
  if [[ -d "$home/.local/bin" ]]; then
    echo "$home/.local/bin"
    return
  fi
  mkdir -p "$home/.local/bin"
  echo "$home/.local/bin"
}

add_to_path() {
  local dir="$1"
  case ":${PATH}:" in
    *":${dir}:"*) return 0 ;;
  esac

  export PATH="${dir}:${PATH}"

  local shell_rc=""
  if [[ -n "${ZSH_VERSION:-}" ]] || [[ "$(basename "$SHELL" 2>/dev/null)" == "zsh" ]]; then
    shell_rc="$HOME/.zshrc"
  elif [[ -n "${BASH_VERSION:-}" ]] || [[ "$(basename "$SHELL" 2>/dev/null)" == "bash" ]]; then
    if [[ -f "$HOME/.bashrc" ]]; then
      shell_rc="$HOME/.bashrc"
    elif [[ -f "$HOME/.bash_profile" ]]; then
      shell_rc="$HOME/.bash_profile"
    fi
  fi

  if [[ -n "$shell_rc" ]]; then
    if ! grep -q "export PATH=\"${dir}:\$PATH\"" "$shell_rc" 2>/dev/null; then
      echo "export PATH=\"${dir}:\$PATH\"" >> "$shell_rc"
      echo "Added ${dir} to PATH in ${shell_rc}"
    fi
  fi

  if command -v fish >/dev/null 2>&1; then
    local fish_config_dir="${XDG_CONFIG_HOME:-$HOME/.config}/fish"
    local fish_conf="$fish_config_dir/conf.d/zaibflow.fish"
    mkdir -p "$fish_config_dir/conf.d"
    if ! grep -q "set -gx PATH ${dir} \$PATH" "$fish_conf" 2>/dev/null; then
      echo "set -gx PATH ${dir} \$PATH" >> "$fish_conf"
      echo "Added ${dir} to PATH for fish in ${fish_conf}"
    fi
  fi
}

create_launchers() {
  local bin_dir="$1"
  local launchers=("zf-kimi" "zf-zai" "zf-or" "zf-local")
  local targets=("kimi" "zai" "openrouter" "ollama")

  for i in "${!launchers[@]}"; do
    local name="${launchers[$i]}"
    local target="${targets[$i]}"
    local path="${bin_dir}/${name}"
    cat > "$path" <<EOF
#!/usr/bin/env bash
exec zaibflow run ${target} "\$@"
EOF
    chmod +x "$path"
  done
}

OS="$(detect_os)"
ARCH="$(detect_arch)"
ASSET="zaibflow_${OS}_${ARCH}.tar.gz"
CHECKSUMS="checksums.txt"
ASSET_PATH="$TMP_DIR/$ASSET"
CHECKSUMS_PATH="$TMP_DIR/$CHECKSUMS"

HOME_DIR="${HOME:-$(eval echo ~${USER:-$(whoami)})}"
INSTALL_DIR="$(get_install_dir "$HOME_DIR")"

echo "Installing ZaibFlow..."
echo "Detected: ${OS} ${ARCH}"
echo "Install dir: ${INSTALL_DIR}"
echo ""

ensure_claude

echo "Downloading: ${ASSET}"
curl -fsSL "$(download_url "$ASSET")" -o "$ASSET_PATH"
curl -fsSL "$(download_url "$CHECKSUMS")" -o "$CHECKSUMS_PATH" || {
  echo "warning: could not download checksums.txt, skipping verification" >&2
  CHECKSUMS_PATH=""
}

if [[ -n "$CHECKSUMS_PATH" ]]; then
  verify_checksum "$ASSET_PATH" "$CHECKSUMS_PATH" || {
    echo "warning: checksum verification failed, continuing install" >&2
  }
fi

tar -xzf "$ASSET_PATH" -C "$TMP_DIR"
mkdir -p "$INSTALL_DIR"
cp "$TMP_DIR/zaibflow" "$INSTALL_DIR/zaibflow"
chmod +x "$INSTALL_DIR/zaibflow"

create_launchers "$INSTALL_DIR"
add_to_path "$INSTALL_DIR"

echo ""
echo "Verifying install..."
if command -v zaibflow >/dev/null 2>&1; then
  zaibflow --version
else
  echo "warning: zaibflow not found on PATH in this session. Run: export PATH=\"${INSTALL_DIR}:\$PATH\"" >&2
fi

echo ""
echo "ZaibFlow installed successfully!"
echo ""
echo "Next commands:"
echo "  zaibflow config"
echo "  zaibflow kimi --bp"
echo "  zaibflow zai --bp"
echo "  zf-kimi --bp"
