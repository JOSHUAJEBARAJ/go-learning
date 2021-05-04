package code

import (
	"crypto/sha1"
	"crypto/sha512"
	"testing"
)

func BenchmarkSHA1(b *testing.B) {

	data := []byte("Hello My name is Joshua")
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		sha1.Sum(data)

	}
}

func BenchmarkSHA512(b *testing.B) {

	data := []byte("Hello My name is Joshua")
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		sha512.Sum(data)
	}
}
