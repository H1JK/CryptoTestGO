#!/bin/bash

export LANG=en_US.UTF-8
export LANGUAGE=en_US.UTF-8

os="UNKNOWN"
osArchitecture="arm-7"

function checkOS() {
  case $(uname -s) in
  Linux*) os="linux" ;;
  Darwin*) os="darwin" ;;
  CYGWIN*) os="windows" ;;
  MINGW*) os="windows" ;;
  *) os="UNKNOWN:${unameOut}" ;;
  esac
}

checkOS

function checkArchitecture() {
  # https://stackoverflow.com/questions/48678152/how-to-detect-386-amd64-arm-or-arm64-os-architecture-via-shell-bash
  case $(uname -m) in
  'i386' | 'i686')
    osArchitecture='386-'
    ;;
  'amd64' | 'x86_64')
    osArchitecture='amd64-'
    ;;
  'armv5tel')
    osArchitecture='arm-5'
    ;;
  'armv6l')
    osArchitecture='arm-6'
    grep Features /proc/cpuinfo | grep -qw 'vfp' || osArchitecture='arm-5'
    ;;
  'armv7' | 'armv7l')
    osArchitecture='arm-7'
    grep Features /proc/cpuinfo | grep -qw 'vfp' || osArchitecture='arm-5'
    ;;
  'armv8' | 'aarch64')
    osArchitecture='arm64-'
    ;;
  'mips')
    osArchitecture='mips-'
    ;;
  'mipsle')
    osArchitecture='mipsle-'
    ;;
  'mips64')
    osArchitecture='mips64-'
    lscpu | grep -q "Little Endian" && osArchitecture='mips64le-'
    ;;
  'mips64le')
    osArchitecture='mips64le-'
    ;;
  'ppc64')
    osArchitecture='ppc64-'
    ;;
  'ppc64le')
    osArchitecture='ppc64le-'
    ;;
  'riscv64')
    osArchitecture='riscv64-'
    ;;
  's390x')
    osArchitecture='s390x-'
    ;;
  *)
    echo "error: The architecture is not supported."
    exit 1
    ;;
  esac
}

checkArchitecture

echo "System: ${os} | Arch: ${osArchitecture}"
echo "Downloading CryptoTestGO..."

if [ os == "windows" ]; then
  fileSuffix=".exe"
else
  fileSuffix=""
fi

fileName="CryptoTestGO-${os}-${osArchitecture}${fileSuffix}"

wget "https://github.com/H1JK/CryptoTestGO/releases/latest/download/${fileName}" -N

chmod +x "${fileName}"
./${fileName} -"test.cpu" 1 -"test.bench"=.
rm ${fileName}
