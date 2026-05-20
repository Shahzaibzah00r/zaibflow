#!/usr/bin/env node

import { access, chmod, copyFile, mkdir, mkdtemp, rm, writeFile } from 'node:fs/promises';
import { createWriteStream } from 'node:fs';
import os from 'node:os';
import path from 'node:path';
import { pipeline } from 'node:stream/promises';
import { Readable } from 'node:stream';
import { spawnSync } from 'node:child_process';

import extractZip from 'extract-zip';
import tar from 'tar';

const OWNER = 'Shahzaibzah00r';
const REPO = 'zaibflow';
const BIN_NAME = 'zaibflow';
const WINDOWS_LAUNCHERS = [
    ['zf-kimi.cmd', ['run', 'kimi']],
    ['zf-or.cmd', ['run', 'openrouter']],
    ['zf-zai.cmd', ['run', 'zai']],
    ['zf-local.cmd', ['run', 'ollama']],
    ['zf-custom.cmd', ['run', 'custom']],
];

const args = process.argv.slice(2);
const platform = detectPlatform();
const arch = detectArch();
const installDir = getInstallDir(platform);
const binaryPath = path.join(installDir, platform === 'win32' ? `${BIN_NAME}.exe` : BIN_NAME);

await ensureInstalled({ platform, arch, installDir, binaryPath });

if (platform === 'win32') {
    await createWindowsLaunchers(installDir, binaryPath);
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
    return path.join(os.homedir(), '.local', 'bin');
}

async function ensureInstalled({ platform: currentPlatform, arch: currentArch, installDir: dir, binaryPath: binPath }) {
    try {
        await access(binPath);
        return;
    } catch {
        // Install from the latest GitHub release when missing.
    }

    await mkdir(dir, { recursive: true });
    const workDir = await mkdtemp(path.join(os.tmpdir(), 'zaibflow-npm-'));
    const assetName = releaseAssetName(currentPlatform, currentArch);
    const archivePath = path.join(workDir, assetName);
    const archiveUrl = releaseAssetUrl(assetName);

    await downloadFile(archiveUrl, archivePath);

    if (currentPlatform === 'win32') {
        await extractZip(archivePath, { dir: workDir });
        await copyFile(path.join(workDir, `${BIN_NAME}.exe`), binPath);
    } else {
        await tar.x({ file: archivePath, cwd: workDir });
        await copyFile(path.join(workDir, BIN_NAME), binPath);
    }

    await chmod(binPath, 0o755).catch(() => { });
    await rm(workDir, { recursive: true, force: true });
}

function releaseAssetName(currentPlatform, currentArch) {
    if (currentPlatform === 'win32') {
        return `${BIN_NAME}_${currentPlatform}_${currentArch}.zip`;
    }
    return `${BIN_NAME}_${currentPlatform}_${currentArch}.tar.gz`;
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
        throw new Error(`download failed: ${response.status} ${response.statusText}`);
    }
    if (!response.body) {
        throw new Error('download failed: empty response body');
    }
    await mkdir(path.dirname(destination), { recursive: true });
    await pipeline(Readable.fromWeb(response.body), createWriteStream(destination));
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

async function writeTextFile(filePath, content) {
    await mkdir(path.dirname(filePath), { recursive: true });
    await rm(filePath, { force: true }).catch(() => { });
    await writeFile(filePath, content, 'utf8');
}

function printInstallSummary(installDir) {
    const lines = [
        `ZaibFlow installed to: ${installDir}`,
        'Next: zaibflow init',
        'Examples:',
        '  zaibflow run kimi',
        '  zaibflow run openrouter <alias>',
        '  zaibflow run ollama',
    ];
    console.log(lines.join('\n'));
}
