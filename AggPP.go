package main

import (
	"crypto/elliptic"
	"crypto/sha256"
	"io"
	"math/big"
)

var ec = elliptic.P256()

// Bunz'20 functions

func bunz_computeZ_sj(a []*big.Int, r *big.Int) (x, y *big.Int) {
	x = big.NewInt(0)
	y = big.NewInt(0)
	n := len(a)

	for i := 0; i < n; i++ {
		scalar := new(big.Int).Exp(r, big.NewInt(int64(2*i)), ec.Params().N)
		scalar.Mul(a[i], scalar)
		scalar.Mod(scalar, ec.Params().N)
		tmpX, tmpY := ec.ScalarBaseMult(scalar.Bytes())
		x, y = ec.Add(x, y, tmpX, tmpY)
	}
	return
}

func Bunz20(a []*big.Int, r *big.Int, ell int) (zX, zY []*big.Int) {
	zX = make([]*big.Int, ell)
	zY = make([]*big.Int, ell)
	for i := 0; i < ell; i++ {
		zX[i], zY[i] = bunz_computeZ_sj(a, r)
	}
	return
}

// Snarkpack functions

func snarkpack_computeZ_sj(a []*big.Int, r *big.Int) (x, y *big.Int) {
	x = big.NewInt(0)
	y = big.NewInt(0)
	scalar := big.NewInt(0)
	n := len(a)

	for i := 0; i < n; i++ {
		tmp := new(big.Int).Exp(r, big.NewInt(int64(i)), ec.Params().N)
		scalar.Add(tmp, scalar)
		scalar.Mod(scalar, ec.Params().N)
	}

	x, y = ec.ScalarBaseMult(scalar.Bytes())
	return
}

func Snackpack(a []*big.Int, r *big.Int, ell int) (zX, zY []*big.Int) {
	zX = make([]*big.Int, ell)
	zY = make([]*big.Int, ell)
	for i := 0; i < ell; i++ {
		zX[i], zY[i] = snarkpack_computeZ_sj(a, r)
	}
	return
}

// aPlonk functions

func aplonk_computePI(a []*big.Int, r *big.Int) (pi *big.Int) {
	pi = big.NewInt(0)
	l := len(a)

	for i := 0; i < l; i++ {
		tmp := new(big.Int).Exp(r, big.NewInt(int64(i)), ec.Params().N)
		pi.Add(tmp, pi)
		pi.Mod(pi, ec.Params().N)
	}
	return
}

func APlonk(a []*big.Int, r, alpha *big.Int, n int) (b *big.Int) {
	b = big.NewInt(0)
	for i := 0; i < n; i++ {
		pi := aplonk_computePI(a, r)
		tmp := new(big.Int).Exp(alpha, big.NewInt(int64(i)), ec.Params().N)
		tmp.Mul(pi, tmp)
		b.Add(tmp, b)
	}
	return
}

// our functions

func Snarkfold(a []*big.Int, n int) []byte {
	l := len(a)
	h := sha256.New()
	for i := 0; i < n; i++ {
		for j := 0; j < l; j++ {
			h.Write(a[j].Bytes())
		}
	}

	return h.Sum(nil)
}

// helper functions

func GenerateRandomF(rand io.Reader) (r *big.Int) {
	N := ec.Params().N
	bitSize := N.BitLen()
	byteLen := (bitSize + 7) / 8
	rbyte := make([]byte, byteLen)
	_, _ = io.ReadFull(rand, rbyte)
	r = new(big.Int).SetBytes(rbyte)
	r = r.Mod(r, N)
	return
}
