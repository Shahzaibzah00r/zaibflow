#!/usr/bin/env node

import { access, chmod, copyFile, mkdir, mkdtemp, readFile, rm, stat, writeFile } from 'node:fs/promises';
import { createWriteStream } from 'node:fs';
import os from 'node:os';
import path from 'node:path';
import { pipeline } from 'node:stream/promises';
import { Readable } from 'node:stream';
import { spawnSync } from 'node:child_process';

import extractZip from 'extract-zip';
import { extract as tarExtract } from 'tar';

const OWNER = 'Shahzaibzah00r';
const REPO = 'zaibflow';
const BIN_NAME = 'zaibflow';
const CACHE_TTL_MS = 24 * 60 * 60 * 1000; // 24 hours

const WINDOWS_LAUNCHERS = [
    ['zf.cmd', []],
    ['zf-kimi.cmd', ['run', 'kimi']],
    ['zf-or.cmd', ['run', 'openrouter']],
    ['zf-zai.cmd', ['run', 'zai']],
    ['zf-local.cmd', ['run', 'ollama']],
    ['zf-custom.cmd', ['run', 'custom']],
];

const UNIX_LAUNCHERS = [
    ['zf', []],
    ['zf-kimi', ['run', 'kimi']],
    ['zf-or', ['run', 'openrouter']],
    ['zf-zai', ['run', 'zai']],
    ['zf-local', ['run', 'ollama']],
    ['zf-custom', ['run', 'custom']],
];

const DEBUG = process.env.ZAIBFLOW_DEBUG === '1';

function debug(...args) {
    if (DEBUG) {
        console.error('[zaibflow npm]', ...args);
    }
}

const args = process.argv.slice(2);
const platform = detectPlatform();
const arch = detectArch();
const installDir = getInstallDir(platform);
const binaryPath = path.join(installDir, platform === 'win32' ? `${BIN_NAME}.exe` : BIN_NAME);

const { installed, updated } = await ensureInstalled({ platform, arch, installDir, binaryPath });

if (installed || updated) {
    debug('Binary was installed/updated; recreating launchers');
    if (platform === 'win32') {
        await createWindowsLaunchers(installDir, binaryPath);
    } else {
        await createUnixLaunchers(installDir);
    }
}

if (args.length === 0) {
    printInstallSummary(installDir);
    process.exit(0);
}

const result = spawnSync(binaryPath, args, { stdio: 'inherit' });
if (result.error) {
    throw result.error;
}
process.exit(result.status ?? 0);

function detectPlatform() {
    if (process.platform === 'win32' || process.platform === 'darwin' || process.platform === 'linux') {
        return process.platform;
    }
    throw new Error(`unsupported platform: ${process.platform}`);
}

function detectArch() {
    if (process.arch === 'x64') {
        return 'amd64';
    }
    if (process.arch === 'arm64') {
        return 'arm64';
    }
    throw new Error(`unsupported architecture: ${process.arch}`);
}

function getInstallDir(currentPlatform) {
    if (currentPlatform === 'win32') {
        const localAppData = process.env.LOCALAPPDATA || path.join(os.homedir(), 'AppData', 'Local');
        return path.join(localAppData, 'Programs', 'ZaibFlow', 'bin');
    }
    const home = os.homedir();
    if (currentPlatform === 'darwin') {
        return path.join(home, 'bin');
    }
    return path.join(home, '.local', 'bin');
}

async function ensureInstalled({ platform: currentPlatform, arch: currentArch, installDir: dir, binaryPath: binPath }) {
    let needsInstall = false;
    let localVersion = null;
    let binaryExists = false;

    try {
        const info = await stat(binPath);
        binaryExists = info.isFile() && info.size > 0;
        debug('Binary exists:', binaryExists, 'size:', info.size);
    } catch {
        binaryExists = false;
        debug('Binary does not exist');
    }

    if (binaryExists) {
        const versionResult = spawnSync(binPath, ['--version'], { encoding: 'utf8', timeout: 5000 });
        const version = (versionResult.stdout || versionResult.stderr || '').trim();
        debug('Local version output:', version);
        // Detect old Clother-based binary (v3.0.9) or any unrecognized version
        if (version.includes('3.0.9') || !version.includes('ZaibFlow')) {
            console.log('ZaibFlow binary is outdated or unrecognized. Re-downloading...');
            needsInstall = true;
        } else {
            localVersion = parseVersion(version);
        }
    } else {
        needsInstall = true;
    }

    if (!needsInstall && localVersion) {
        const latestVersion = await getLatestVersion();
        debug('Latest version:', latestVersion ? formatVersion(latestVersion) : 'unknown');
        if (latestVersion && isNewer(latestVersion, localVersion)) {
            console.log(`ZaibFlow update available: v${formatVersion(localVersion)} -> v${formatVersion(latestVersion)}. Downloading...`);
            needsInstall = true;
        }
    }

    if (!needsInstall) {
        return { installed: false, updated: false };
    }

    await mkdir(dir, { recursive: true });
    const workDir = await mkdtemp(path.join(os.tmpdir(), 'zaibflow-npm-'));
    const assetName = releaseAssetName(currentPlatform, currentArch);
    const archivePath = path.join(workDir, assetName);
    const archiveUrl = releaseAssetUrl(assetName);

    debug('Downloading from:', archiveUrl);

    try {
        await downloadFile(archiveUrl, archivePath);
    } catch (err) {
        console.error(`Failed to download ${assetName} from ${archiveUrl}`);
        console.error(err.message);
        throw err;
    }

    if (currentPlatform === 'win32') {
        await extractZip(archivePath, { dir: workDir });
        await copyFile(path.join(workDir, `${BIN_NAME}.exe`), binPath);
    } else {
        await tarExtract({ file: archivePath, cwd: workDir });
        await copyFile(path.join(workDir, BIN_NAME), binPath);
    }

    await chmod(binPath, 0o755).catch(() => { });
    await rm(workDir, { recursive: true, force: true });

    return { installed: true, updated: binaryExists };
}

function releaseAssetName(currentPlatform, currentArch) {
    const platform = currentPlatform === 'win32' ? 'windows' : currentPlatform;
    if (currentPlatform === 'win32') {
        return `${BIN_NAME}_${platform}_${currentArch}.zip`;
    }
    return `${BIN_NAME}_${platform}_${currentArch}.tar.gz`;
}

function releaseAssetUrl(assetName) {
    const baseUrl = (process.env.ZAIBFLOW_RELEASE_BASE_URL || '').trim().replace(/\/$/, '');
    if (baseUrl) {
        return `${baseUrl}/${assetName}`;
    }
    const version = (process.env.ZAIBFLOW_VERSION || 'latest').trim();
    if (version && version !== 'latest') {
        return `https://github.com/${OWNER}/${REPO}/releases/download/${version}/${assetName}`;
    }
    return `https://github.com/${OWNER}/${REPO}/releases/latest/download/${assetName}`;
}

async function downloadFile(url, destination) {
    const response = await fetch(url, {
        headers: {
            'User-Agent': 'zaibflow-npm-installer',
        },
    });
    if (!response.ok) {
        throw new Error(`download failed: ${response.status} ${response.statusText} for ${url}`);
    }
    if (!response.body) {
        throw new Error('download failed: empty response body');
    }
    await mkdir(path.dirname(destination), { recursive: true });
    await pipeline(Readable.fromWeb(response.body), createWriteStream(destination));
}

async function getLatestVersion() {
    const cacheDir = path.join(os.tmpdir(), 'zaibflow-npm-cache');
    const cacheFile = path.join(cacheDir, 'latest-version.json');

    try {
        const cache = JSON.parse(await readFile(cacheFile, 'utf8'));
        if (cache.timestamp && Date.now() - cache.timestamp < CACHE_TTL_MS) {
            return parseVersion(cache.version);
        }
    } catch {
        // cache missing or invalid
    }

    try {
        const response = await fetch(`https://api.github.com/repos/${OWNER}/${REPO}/releases/latest`, {
            headers: {
                'User-Agent': 'zaibflow-npm-installer',
                'Accept': 'application/vnd.github+json',
            },
            signal: AbortSignal.timeout(5000),
        });
        if (!response.ok) {
            return null;
        }
        const data = await response.json();
        const version = data.tag_name?.replace(/^v/, '');
        if (version) {
            await mkdir(cacheDir, { recursive: true });
            await writeFile(cacheFile, JSON.stringify({ version, timestamp: Date.now() }), 'utf8');
            return parseVersion(version);
        }
        return null;
    } catch {
        return null;
    }
}

function parseVersion(versionString) {
    const match = versionString.match(/v?(\d+)\.(\d+)\.(\d+)/);
    if (!match) return null;
    return {
        major: parseInt(match[1], 10),
        minor: parseInt(match[2], 10),
        patch: parseInt(match[3], 10),
    };
}

function formatVersion(v) {
    if (!v) return 'unknown';
    return `${v.major}.${v.minor}.${v.patch}`;
}

function isNewer(latest, current) {
    if (!latest || !current) return false;
    if (latest.major !== current.major) return latest.major > current.major;
    if (latest.minor !== current.minor) return latest.minor > current.minor;
    return latest.patch > current.patch;
}

async function createWindowsLaunchers(installDir, binaryPath) {
    for (const [fileName, commandArgs] of WINDOWS_LAUNCHERS) {
        const content = [
            '@echo off',
            'setlocal',
            `"%~dp0${path.basename(binaryPath)}" ${commandArgs.join(' ')} %*`,
            'exit /b %errorlevel%',
            '',
        ].join('\r\n');
        await writeTextFile(path.join(installDir, fileName), content);
    }
}

async function createUnixLaunchers(installDir) {
    for (const [fileName, commandArgs] of UNIX_LAUNCHERS) {
        const content = `#!/usr/bin/env bash\nexec zaibflow ${commandArgs.join(' ')} "$@"\n`;
        const filePath = path.join(installDir, fileName);
        await writeTextFile(filePath, content);
        await chmod(filePath, 0o755).catch(() => { });
    }
}

async function writeTextFile(filePath, content) {
    await mkdir(path.dirname(filePath), { recursive: true });
    await rm(filePath, { force: true }).catch(() => { });
    await writeFile(filePath, content, 'utf8');
}

function printInstallSummary(installDir) {
    const lines = [
        `ZaibFlow installed to: ${installDir}`,
        '',
        'Next commands:',
        '  zaibflow config',
        '  zaibflow kimi --bp',
        '  zaibflow zai --bp',
        '  zf-kimi --bp',
        '',
        'Tip: If zaibflow is not found, restart your terminal so PATH changes take effect.',
    ];
    console.log(lines.join('\n'));
}
