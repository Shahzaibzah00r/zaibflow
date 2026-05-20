package launchers

import (
	"encoding/json"
	"os"
	"path/filepath"
	goruntime "runtime"
	"sort"
	"strings"

	"github.com/Shahzaibzah00r/zaibflow/internal/config"
	"github.com/Shahzaibzah00r/zaibflow/internal/profiles"
	"github.com/Shahzaibzah00r/zaibflow/internal/providers"
)

type Manifest struct {
	Launchers []string `json:"launchers"`
}

// Sync installs the ZaibFlow binary and provider launchers into paths.BinDir.
//
// When skipCopy is false (normal install), the binary at execPath is copied to
// paths.BinDir/zaibflow and symlinks or launcher shims are created relative to it.
//
// When skipCopy is true (Homebrew install), no binary is copied; symlinks are
// created as absolute references to execPath so that a Homebrew-managed binary
// upgrade is reflected automatically without running `zaibflow install` again.
func Sync(execPath string, paths config.Paths, catalog providers.Catalog, cfg *config.File, skipCopy bool) error {
	if err := paths.EnsureBaseDirs(); err != nil {
		return err
	}

	previous, _ := LoadManifest(paths.ManifestFile)
	desired := map[string]struct{}{}
	for _, target := range profiles.All(catalog, cfg) {
		// Under Homebrew the formula already installs static provider launchers in
		// the Homebrew prefix, and zf-or / zf-custom cover dynamic
		// providers via gateway invocation. Skip individual dynamic symlinks to
		// keep ~/bin clean for Homebrew users.
		if skipCopy && isDynamicProfile(target.Profile, cfg) {
			continue
		}
		desired[LauncherName(target.Profile)] = struct{}{}
	}
	// Always create gateway symlinks regardless of install method or whether
	// any dynamic providers are configured. The isDynamicProfile skip above
	// only applies to per-alias/per-provider symlinks, never to these gateways.
	desired["zf-or"] = struct{}{}
	desired["zf-custom"] = struct{}{}

	if goruntime.GOOS == "windows" {
		destBinary := filepath.Join(paths.BinDir, binaryName())
		if !skipCopy {
			if err := copyExecutable(execPath, destBinary); err != nil {
				return err
			}
		}
		for _, old := range previous.Launchers {
			if _, ok := desired[old]; ok {
				continue
			}
			_ = os.Remove(filepath.Join(paths.BinDir, old))
		}
		var launchers []string
		for name := range desired {
			launchers = append(launchers, name)
		}
		sort.Strings(launchers)
		for _, name := range launchers {
			path := filepath.Join(paths.BinDir, name+".cmd")
			_ = os.Remove(path)
			if err := writeAtomic(path, windowsLauncher(binaryName(), launcherCommand(name, cfg), name), 0o755); err != nil {
				return err
			}
		}
		return SaveManifest(paths.ManifestFile, Manifest{Launchers: launchers})
	}

	symlinkTarget := binaryName() // relative — works when binary lives in the same dir
	if skipCopy {
		symlinkTarget = execPath // absolute — points directly to the Homebrew binary
	} else {
		destBinary := filepath.Join(paths.BinDir, binaryName())
		if err := copyExecutable(execPath, destBinary); err != nil {
			return err
		}
	}

	for _, old := range previous.Launchers {
		if _, ok := desired[old]; ok {
			continue
		}
		_ = os.Remove(filepath.Join(paths.BinDir, old))
	}

	var launchers []string
	for name := range desired {
		launchers = append(launchers, name)
	}
	sort.Strings(launchers)
	for _, name := range launchers {
		link := filepath.Join(paths.BinDir, name)
		_ = os.Remove(link)
		if err := os.Symlink(symlinkTarget, link); err != nil {
			return err
		}
	}
	claudeShim := filepath.Join(paths.BinDir, "claude")
	_ = os.Remove(claudeShim)
	if err := os.Symlink(symlinkTarget, claudeShim); err != nil {
		return err
	}
	return SaveManifest(paths.ManifestFile, Manifest{Launchers: launchers})
}

func LoadManifest(path string) (Manifest, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Manifest{}, err
	}
	var manifest Manifest
	if err := json.Unmarshal(data, &manifest); err != nil {
		return Manifest{}, err
	}
	return manifest, nil
}

func SaveManifest(path string, manifest Manifest) error {
	data, err := json.MarshalIndent(manifest, "", "  ")
	if err != nil {
		return err
	}
	data = append(data, '\n')
	return writeAtomic(path, data, 0o644)
}

// isDynamicProfile reports whether a profile is user-defined (OpenRouter alias
// or custom provider) rather than a catalog-builtin static provider.
func isDynamicProfile(profile string, cfg *config.File) bool {
	if strings.HasPrefix(profile, "or-") {
		return true
	}
	if cfg == nil {
		return false
	}
	_, isCustom := cfg.CustomProviders[profile]
	return isCustom
}

func copyExecutable(src, dst string) error {
	if err := os.MkdirAll(filepath.Dir(dst), 0o755); err != nil {
		return err
	}
	data, err := os.ReadFile(src)
	if err != nil {
		return err
	}
	return writeAtomic(dst, data, 0o755)
}

func writeAtomic(path string, data []byte, mode os.FileMode) error {
	tmp, err := os.CreateTemp(filepath.Dir(path), ".launcher-*")
	if err != nil {
		return err
	}
	tmpPath := tmp.Name()
	defer os.Remove(tmpPath)
	if _, err := tmp.Write(data); err != nil {
		tmp.Close()
		return err
	}
	if err := tmp.Chmod(mode); err != nil {
		tmp.Close()
		return err
	}
	if err := tmp.Close(); err != nil {
		return err
	}
	return os.Rename(tmpPath, path)
}

func binaryName() string {
	if goruntime.GOOS == "windows" {
		return "zaibflow.exe"
	}
	return "zaibflow"
}

func launcherCommand(name string, cfg *config.File) []string {
	if name == "zf-or" {
		return []string{"run", "openrouter"}
	}
	if name == "zf-custom" {
		return []string{"run", "custom"}
	}
	profile := strings.TrimPrefix(name, "zf-")
	if profile == "local" {
		return []string{"run", "ollama"}
	}
	if strings.HasPrefix(profile, "or-") {
		return []string{"run", "openrouter", strings.TrimPrefix(profile, "or-")}
	}
	if cfg != nil {
		if _, ok := cfg.CustomProviders[profile]; ok {
			return []string{"run", "custom", profile}
		}
	}
	return []string{"run", profile}
}

func windowsLauncher(binary string, args []string, displayName string) []byte {
	var builder strings.Builder
	builder.WriteString("@echo off\r\n")
	builder.WriteString("setlocal\r\n")
	builder.WriteString("\"%~dp0")
	builder.WriteString(binary)
	builder.WriteString("\"")
	for _, arg := range args {
		builder.WriteString(" ")
		builder.WriteString(arg)
	}
	builder.WriteString(" %*\r\n")
	builder.WriteString("exit /b %errorlevel%\r\n")
	return []byte(builder.String())
}
