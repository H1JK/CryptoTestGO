package main

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"runtime"
	"strconv"
	"testing"

	"golang.org/x/crypto/chacha20poly1305"
)

func benchmarkAESGCMEncrypt(b *testing.B, buf []byte, keySize int) {
	b.ReportAllocs()
	b.SetBytes(int64(len(buf)))

	var key = make([]byte, keySize)
	var nonce [12]byte
	var ad [13]byte
	aes, _ := aes.NewCipher(key[:])
	aesgcm, _ := cipher.NewGCM(aes)
	var out []byte

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		out = aesgcm.Seal(out[:0], nonce[:], buf, ad[:])
	}
}

func benchmarkAESGCMDecrypt(b *testing.B, buf []byte, keySize int) {
	b.ReportAllocs()
	b.SetBytes(int64(len(buf)))

	var key = make([]byte, keySize)
	var nonce [12]byte
	var ad [13]byte
	aes, _ := aes.NewCipher(key[:])
	aesgcm, _ := cipher.NewGCM(aes)
	var out []byte

	ct := aesgcm.Seal(nil, nonce[:], buf[:], ad[:])

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		out, _ = aesgcm.Open(out[:0], nonce[:], ct, ad[:])
	}
}

func benchmarkAESGCM(b *testing.B) {
	for _, length := range []int{64, 1350, 8 * 1024} {
		b.Run("Decrypt-128-"+strconv.Itoa(length), func(b *testing.B) {
			benchmarkAESGCMDecrypt(b, make([]byte, length), 128/8)
		})
	}
	for _, length := range []int{64, 1350, 8 * 1024} {
		b.Run("Decrypt-256-"+strconv.Itoa(length), func(b *testing.B) {
			benchmarkAESGCMDecrypt(b, make([]byte, length), 256/8)
		})
	}
	for _, length := range []int{64, 1350, 8 * 1024} {
		b.Run("Encrypt-128-"+strconv.Itoa(length), func(b *testing.B) {
			benchmarkAESGCMEncrypt(b, make([]byte, length), 128/8)
		})
	}
	for _, length := range []int{64, 1350, 8 * 1024} {
		b.Run("Encrypt-256-"+strconv.Itoa(length), func(b *testing.B) {
			benchmarkAESGCMEncrypt(b, make([]byte, length), 256/8)
		})
	}
}

func benchamarkChaCha20Poly1305Encrypt(b *testing.B, buf []byte, nonceSize int) {
	b.ReportAllocs()
	b.SetBytes(int64(len(buf)))

	var key [32]byte
	var nonce = make([]byte, nonceSize)
	var ad [13]byte
	var out []byte

	var aead cipher.AEAD
	switch len(nonce) {
	case chacha20poly1305.NonceSize:
		aead, _ = chacha20poly1305.New(key[:])
	case chacha20poly1305.NonceSizeX:
		aead, _ = chacha20poly1305.NewX(key[:])
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		out = aead.Seal(out[:0], nonce[:], buf[:], ad[:])
	}
}

func benchamarkChaCha20Poly1305Decrypt(b *testing.B, buf []byte, nonceSize int) {
	b.ReportAllocs()
	b.SetBytes(int64(len(buf)))

	var key [32]byte
	var nonce = make([]byte, nonceSize)
	var ad [13]byte
	var ct []byte
	var out []byte

	var aead cipher.AEAD
	switch len(nonce) {
	case chacha20poly1305.NonceSize:
		aead, _ = chacha20poly1305.New(key[:])
	case chacha20poly1305.NonceSizeX:
		aead, _ = chacha20poly1305.NewX(key[:])
	}
	ct = aead.Seal(ct[:0], nonce[:], buf[:], ad[:])

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		out, _ = aead.Open(out[:0], nonce[:], ct[:], ad[:])
	}
}

func benchmarkChacha20Poly1305(b *testing.B) {
	for _, length := range []int{64, 1350, 8 * 1024} {
		b.Run("Decrypt-"+strconv.Itoa(length), func(b *testing.B) {
			benchamarkChaCha20Poly1305Decrypt(b, make([]byte, length), chacha20poly1305.NonceSize)
		})
	}
	for _, length := range []int{64, 1350, 8 * 1024} {
		b.Run("Decrypt-"+strconv.Itoa(length)+"-X", func(b *testing.B) {
			benchamarkChaCha20Poly1305Decrypt(b, make([]byte, length), chacha20poly1305.NonceSizeX)
		})
	}
	for _, length := range []int{64, 1350, 8 * 1024} {
		b.Run("Encrypt-"+strconv.Itoa(length), func(b *testing.B) {
			benchamarkChaCha20Poly1305Encrypt(b, make([]byte, length), chacha20poly1305.NonceSize)
		})
	}
	for _, length := range []int{64, 1350, 8 * 1024} {
		b.Run("Encrypt-"+strconv.Itoa(length)+"-X", func(b *testing.B) {
			benchamarkChaCha20Poly1305Encrypt(b, make([]byte, length), chacha20poly1305.NonceSizeX)
		})
	}
}

func BenchmarkCrypto(b *testing.B) {
	fmt.Printf("CryptoTestGO %s (%s, %s/%s)\n", Version, runtime.Version(), runtime.GOOS, runtime.GOARCH)
	fmt.Println("GitHub: https://github.com/H1JK/CryptoTestGO")
	old, new := GetSecurityAutoType()
	fmt.Printf("VMess AUTO will choose: Xray/V2Fly 5.0.8-: %s, V2Fly 5.0.8+:%s\n", old, new)
	b.Run("AES-GCM", benchmarkAESGCM)
	b.Run("ChaCha20-Poly1305", benchmarkChacha20Poly1305)
}
