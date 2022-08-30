package main

import (
	"fmt"
	"runtime"

	"golang.org/x/sys/cpu"
)

func main() {
	fmt.Printf("CryptoTestGO %s (%s, %s ,%s)\n", Version, runtime.Version(), runtime.GOOS, runtime.GOARCH)
	fmt.Println("GitHub: https://github.com/H1JK/CryptoTestGO")
	fmt.Println("\nIllegal usage. If you want to run from code, use `go test -cpu 1 -benchmem -bench=BenchmarkCrypto ./...`")
}

func GetSecurityAutoType() (old string, new string) {
	if runtime.GOARCH == "amd64" || runtime.GOARCH == "arm64" || runtime.GOARCH == "s390x" {
		old = "AES-128-GCM"
	} else {
		old = "ChaCha20-Poly1305"
	}
	new = "ChaCha20-Poly1305"
	switch runtime.GOARCH {
	case "amd64":
		if cpu.X86.HasAES && cpu.X86.HasPCLMULQDQ {
			new = "AES-128-GCM"
		}
	case "arm64":
		if cpu.ARM64.HasAES && cpu.ARM64.HasPMULL {
			new = "AES-128-GCM"
		}
	case "s390x":
		if cpu.S390X.HasAES && cpu.S390X.HasAESCBC && cpu.S390X.HasAESCTR &&
			(cpu.S390X.HasGHASH || cpu.S390X.HasAESGCM) {
			new = "AES-128-GCM"
		}
	}
	return
}
