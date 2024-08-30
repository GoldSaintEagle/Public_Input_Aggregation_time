package main

import (
	"crypto/rand"
	"math/big"
	"testing"
)

var ell = 200
var n = 100

func BenchmarkBunz20(b *testing.B) {

	r := GenerateRandomF(rand.Reader)
	a := make([]*big.Int, n)

	for i := 0; i < n; i++ {
		a[i] = GenerateRandomF(rand.Reader)
	}

	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, _ = Bunz20(a, r, ell)
		}
	})
}

func BenchmarkSnackpack(b *testing.B) {

	r := GenerateRandomF(rand.Reader)
	a := make([]*big.Int, n)

	for i := 0; i < n; i++ {
		a[i] = GenerateRandomF(rand.Reader)
	}

	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, _ = Snackpack(a, r, ell)
		}
	})
}

func BenchmarkAPlonk(b *testing.B) {

	r := GenerateRandomF(rand.Reader)
	alpha := GenerateRandomF(rand.Reader)
	a := make([]*big.Int, ell)

	for i := 0; i < ell; i++ {
		a[i] = GenerateRandomF(rand.Reader)
	}

	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = APlonk(a, r, alpha, n)
		}
	})
}

func BenchmarkSnarkfold(b *testing.B) {
	a := make([]*big.Int, ell)

	for i := 0; i < ell; i++ {
		a[i] = GenerateRandomF(rand.Reader)
	}

	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = Snarkfold(a, n)
		}
	})
}
