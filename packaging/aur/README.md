# AUR Packaging

This directory contains AUR packaging assets for OwnCord.

## Packages

- owncord-server-git: development package that builds the Go server from the latest git source.
- owncord-client-git: development package that builds the Tauri desktop client from the latest git source.

Files live in:

- packaging/aur/owncord-server-git/PKGBUILD
- packaging/aur/owncord-server-git/.SRCINFO
- packaging/aur/owncord-client-git/PKGBUILD
- packaging/aur/owncord-client-git/.SRCINFO

## Local build test

```bash
cd packaging/aur/owncord-server-git
makepkg -si

cd ../owncord-client-git
makepkg -si
```

## Publish to AUR

### Server package

```bash
git clone ssh://aur@aur.archlinux.org/owncord-server-git.git
cd owncord-server-git
cp -r /path/to/OwnCord/packaging/aur/owncord-server-git/* .
makepkg --printsrcinfo > .SRCINFO
git add PKGBUILD .SRCINFO owncord-server.service owncord-server.sysusers owncord-server.tmpfiles owncord-server-git.install
git commit -m "Initial package"
git push
```

### Client package

```bash
git clone ssh://aur@aur.archlinux.org/owncord-client-git.git
cd owncord-client-git
cp -r /path/to/OwnCord/packaging/aur/owncord-client-git/* .
makepkg --printsrcinfo > .SRCINFO
git add PKGBUILD .SRCINFO owncord-client.desktop
git commit -m "Initial package"
git push
```

## Runtime

After install:

```bash
sudo systemctl enable --now owncord-server.service
```

The service runs as user `owncord` with working directory `/var/lib/owncord`.
The first start generates `config.yaml` and `data/` automatically.

Client app entry point after install: `owncord-client`.
