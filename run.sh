#!/bin/bash

export LANG=en_US.UTF-8
export LANGUAGE=en_US.UTF-8

os="UNKNOWN"
osArchitecture="arm-7"

function checkOS() {
  case $(uname -s) in
    Linux*)     os="linux";;
    Darwin*)    os="darwin";;
    CYGWIN*)    os="windows";;
    MINGW*)     os="windows";;
    *)          os="UNKNOWN:${unameOut}"
  esac
}

checkOS

function checkArchitecture(){
	# https://stackoverflow.com/questions/48678152/how-to-detect-386-amd64-arm-or-arm64-os-architecture-via-shell-bash

	case $(uname -m) in
		i386)   osArchitecture="386-" ;;
		i686)   osArchitecture="386-" ;;
		x86_64) osArchitecture="amd64-" ;;
		arm)    dpkg --print-architecture | grep -q "arm64" && osArchitecture="arm64-" || osArchitecture="arm-7" ;;
		aarch64)    dpkg --print-architecture | grep -q "arm64" && osArchitecture="arm64-" || osArchitecture="arm-7" ;;
		* )     osArchitecture="arm-7" ;;
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

wget "https://download.fastgit.org/H1JK/CryptoTestGO/releases/latest/download/${fileName}" -N

chmod +x "${fileName}"
./${fileName} -"test.cpu" 1 -"test.bench"=.
rm ${fileName}
