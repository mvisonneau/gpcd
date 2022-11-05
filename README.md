# gpcd - GoPro Cloud Downloader

[![PkgGoDev](https://pkg.go.dev/badge/github.com/mvisonneau/gpcd)](https://pkg.go.dev/mod/github.com/mvisonneau/gpcd)
[![Go Report Card](https://goreportcard.com/badge/github.com/mvisonneau/gpcd)](https://goreportcard.com/report/github.com/mvisonneau/gpcd)
[![test](https://github.com/mvisonneau/gpcd/actions/workflows/test.yml/badge.svg)](https://github.com/mvisonneau/gpcd/actions/workflows/test.yml)
[![Coverage Status](https://coveralls.io/repos/github/mvisonneau/gpcd/badge.svg?branch=main)](https://coveralls.io/github/mvisonneau/gpcd?branch=main)
[![release](https://github.com/mvisonneau/gpcd/actions/workflows/release.yml/badge.svg)](https://github.com/mvisonneau/gpcd/actions/workflows/release.yml)
[![gpcd](https://snapcraft.io/mvisonneau-gpcd/badge.svg)](https://snapcraft.io/mvisonneau-gpcd)

Hacky command-line tool to go beyond their current <= 25 media download limitation and to download files concurrently.

## Example

Download all medias captured since a particular date:

```
# List & review them
~$ gpcd --from "2022-10-24T00:00:00Z" list
nrWxYZXMOld7E | 2022-10-24 21:00:09 +0000 UTC | GX003942.MP4 - Video (2988p)
dMepexYZsAsOa | 2022-10-24 07:10:51 +0000 UTC | GOPR0311.JPG - Photo (27127296)
V2EWErxjHdjmJ | 2022-10-24 07:10:24 +0000 UTC | GOPR9890.JPG - Photo (27127296)
7azxjdKsjBv83 | 2022-10-24 07:09:36 +0000 UTC | GX010400.MP4 - Video (2988p)

# Download them
~$ gpcd --from "2022-10-24T00:00:00Z" download
nrWxYZXMOld7E | 2022-10-24 21:00:09 +0000 UTC | GX003942.MP4 - Video (2988p)
downloading 100% |██████████████████████████████████████████████████████████████████████| (92/92 MB, 5.3 MB/s)
dMepexYZsAsOa | 2022-10-24 07:10:51 +0000 UTC | GOPR0311.JPG - Photo (27127296)
downloading 100% |██████████████████████████████████████████████████████████████████████| (5.5/5.5 MB, 3.6 MB/s)
V2EWErxjHdjmJ | 2022-10-24 07:10:24 +0000 UTC | GOPR9890.JPG - Photo (27127296)
downloading 100% |██████████████████████████████████████████████████████████████████████| (6.1/6.1 MB, 4.0 MB/s)
7azxjdKsjBv83 | 2022-10-24 07:09:36 +0000 UTC | GX010400.MP4 - Video (2988p)
downloading 100% |██████████████████████████████████████████████████████████████████████| (194/194 MB, 14 MB/s)
```

## Usage

```
~$ gpcd
NAME:
   gpcd - download bulk medias from GoPro Cloud

USAGE:
   gpcd [global options] command [command options] [arguments...]

COMMANDS:
   list      list available medias
   download  download medias
   help, h   Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --api-endpoint value              Go Pro Cloud API endpoint (default: "https://api.gopro.com/media/") [$GPCD_API_ENDPOINT]
   --local-path value                where the medias should be downloaded (default: "./medias") [$GPCD_LOCAL_PATH]
   --bearer-token value              Used to authenticate over your account [$GPCD_BEARER_TOKEN]
   --from value                      filter for medias captured after this date
   --to value                        filter for medias captured before this date
   --max-concurrent-downloads value  (default: 8)
   --help, -h                        show help (default: false)
```

## Install

Have a look onto the [latest release page](https://github.com/mvisonneau/gpcd/releases/latest) and pick your flavor.

Checksums are signed with the [following GPG key](https://keybase.io/mvisonneau/pgp_keys.asc): `C09C A9F7 1C5C 988E 65E3  E5FC ADEA 38ED C46F 25BE`

### Go

```bash
~$ go install github.com/mvisonneau/gpcd/cmd/gpcd@latest
```

### Homebrew

```bash
~$ brew install mvisonneau/tap/gpcd
```

### Docker

```bash
~$ docker run -it --rm docker.io/mvisonneau/gpcd
~$ docker run -it --rm ghcr.io/mvisonneau/gpcd
~$ docker run -it --rm quay.io/mvisonneau/gpcd
```

### Scoop

```bash
~$ scoop bucket add https://github.com/mvisonneau/scoops
~$ scoop install gpcd
```

### Binaries, DEB and RPM packages

For the following ones, you need to know which version you want to install, to fetch the latest available :

```bash
~$ export GPCD_VERSION=$(curl -s "https://api.github.com/repos/mvisonneau/gpcd/releases/latest" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')
```

```bash
# Binary (eg: freebsd/amd64)
~$ wget https://github.com/mvisonneau/gpcd/releases/download/${GPCD_VERSION}/gpcd_${GPCD_VERSION}_freebsd_amd64.tar.gz
~$ tar zxvf gpcd_${GPCD_VERSION}_freebsd_amd64.tar.gz -C /usr/local/bin

# DEB package (eg: linux/386)
~$ wget https://github.com/mvisonneau/gpcd/releases/download/${GPCD_VERSION}/gpcd_${GPCD_VERSION}_linux_386.deb
~$ dpkg -i gpcd_${GPCD_VERSION}_linux_386.deb

# RPM package (eg: linux/arm64)
~$ wget https://github.com/mvisonneau/gpcd/releases/download/${GPCD_VERSION}/gpcd_${GPCD_VERSION}_linux_arm64.rpm
~$ rpm -ivh gpcd_${GPCD_VERSION}_linux_arm64.rpm
```
