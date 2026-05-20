param(
  [string]$Version = "latest"
)

$ErrorActionPreference = 'Stop'
$ProgressPreference = 'SilentlyContinue'

$Owner = 'Shahzaibzah00r'
$Repo = 'zaibflow'
$BinaryName = 'zaibflow.exe'
$BinDir = if ($env:LOCALAPPDATA) {
  Join-Path $env:LOCALAPPDATA 'Programs\ZaibFlow\bin'
} else {
  Join-Path $HOME 'AppData\Local\Programs\ZaibFlow\bin'
}

function Get-Architecture {
  $arch = $env:PROCESSOR_ARCHITECTURE
  if (-not $arch -and $env:PROCESSOR_ARCHITEW6432) {
    $arch = $env:PROCESSOR_ARCHITEW6432
  }
  if (-not $arch) {
    $arch = ''
  }
  switch ($arch.ToLowerInvariant()) {
    'amd64' { return 'amd64' }
    'x86_64' { return 'amd64' }
    'arm64' { return 'arm64' }
    default { throw "unsupported architecture: $arch" }
  }
}

function Get-ReleaseAssetUrl {
  param(
    [string]$Asset
  )

  $baseUrl = $env:ZAIBFLOW_RELEASE_BASE_URL
  if ($baseUrl) {
    return ($baseUrl.TrimEnd('/') + '/' + $Asset)
  }

  $tag = $Version
  if (-not $tag) {
    $tag = 'latest'
  }
  if ($tag -eq 'latest') {
    return "https://github.com/$Owner/$Repo/releases/latest/download/$Asset"
  }
  return "https://github.com/$Owner/$Repo/releases/download/$tag/$Asset"
}

function Ensure-UserPath {
  param(
    [string]$PathEntry
  )

  $userPath = [Environment]::GetEnvironmentVariable('Path', 'User')
  if (-not $userPath) {
    [Environment]::SetEnvironmentVariable('Path', $PathEntry, 'User')
    return
  }

  $entries = $userPath.Split(';') | Where-Object { $_ -and $_.Trim() }
  foreach ($entry in $entries) {
    if ($entry.TrimEnd('\') -ieq $PathEntry.TrimEnd('\')) {
      $env:Path = "$PathEntry;$env:Path"
      return
    }
  }

  [Environment]::SetEnvironmentVariable('Path', ($userPath.TrimEnd(';') + ';' + $PathEntry), 'User')
  $env:Path = "$PathEntry;$env:Path"
}

function Write-CmdLauncher {
  param(
    [string]$FilePath,
    [string[]]$Args
  )

  $content = @(
    '@echo off',
    'setlocal',
    ('"%~dp0' + $BinaryName + '" ' + ($Args -join ' ') + ' %*'),
    'exit /b %errorlevel%',
    ''
  ) -join "`r`n"
  Set-Content -Path $FilePath -Value $content -Encoding Ascii
}

function Ensure-Claude {
  $claude = Get-Command 'claude' -ErrorAction SilentlyContinue
  if ($claude) {
    return
  }

  Write-Host "Claude Code CLI is required."
  Write-Host ""
  Write-Host "Please install Claude Code CLI manually:"
  Write-Host "  See https://docs.anthropic.com/en/docs/agents-and-tools/claude-code/overview"
  Write-Host ""
  Write-Host "After installing Claude Code, re-run this installer."
  exit 1
}

Ensure-Claude

$arch = Get-Architecture
$assetName = "zaibflow_windows_${arch}.zip"
$assetUrl = Get-ReleaseAssetUrl -Asset $assetName

$tempDir = New-Item -ItemType Directory -Path ([System.IO.Path]::Combine([System.IO.Path]::GetTempPath(), "zaibflow-$(Get-Random)"))
$zipPath = Join-Path $tempDir.FullName $assetName
$extractDir = Join-Path $tempDir.FullName 'extract'
New-Item -ItemType Directory -Path $extractDir | Out-Null

Invoke-WebRequest -Uri $assetUrl -OutFile $zipPath
Expand-Archive -Path $zipPath -DestinationPath $extractDir -Force

New-Item -ItemType Directory -Path $BinDir -Force | Out-Null
Copy-Item -Path (Join-Path $extractDir $BinaryName) -Destination (Join-Path $BinDir $BinaryName) -Force

# Main shortcut: zf behaves exactly like zaibflow
$content = @(
  '@echo off',
  'setlocal',
  ('"%~dp0' + $BinaryName + '" %*'),
  'exit /b %errorlevel%',
  ''
) -join "`r`n"
Set-Content -Path (Join-Path $BinDir 'zf.cmd') -Value $content -Encoding Ascii

Write-CmdLauncher -FilePath (Join-Path $BinDir 'zf-kimi.cmd') -Args @('run', 'kimi')
Write-CmdLauncher -FilePath (Join-Path $BinDir 'zf-or.cmd') -Args @('run', 'openrouter')
Write-CmdLauncher -FilePath (Join-Path $BinDir 'zf-zai.cmd') -Args @('run', 'zai')
Write-CmdLauncher -FilePath (Join-Path $BinDir 'zf-local.cmd') -Args @('run', 'ollama')
Write-CmdLauncher -FilePath (Join-Path $BinDir 'zf-custom.cmd') -Args @('run', 'custom')

Ensure-UserPath -PathEntry $BinDir

Write-Host ""
Write-Host "Verifying install..."
$zaibflow = Get-Command 'zaibflow' -ErrorAction SilentlyContinue
if ($zaibflow) {
  & zaibflow --version
} else {
  Write-Warning "zaibflow not found on PATH in this session. Restart your terminal."
}

Write-Host ""
Write-Host "ZaibFlow installed successfully!"
Write-Host ""
Write-Host "Next commands:"
Write-Host "  zaibflow config"
Write-Host "  zaibflow kimi --bp"
Write-Host "  zaibflow zai --bp"
Write-Host "  zf-kimi --bp"
