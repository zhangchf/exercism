package diffiehellman

import (
	"crypto/rand"
	"math/big"
)

//type testCase struct {
//	g    int64    // prime, generator
//	p    *big.Int // prime, modulus
//	a, b *big.Int // private keys
//	A, B *big.Int // public keys
//	s    *big.Int // secret key
//}

var __one = big.NewInt(1)

func PrivateKey(p *big.Int) *big.Int {
	for {
		if a, err := rand.Int(rand.Reader, p); err == nil {
			if a.Cmp(__one) > 0 {
				return a
			}
		}
	}
	return nil
}

func PublicKey(a, p *big.Int, g int64) (A *big.Int) {
	return new(big.Int).Exp(big.NewInt(g), a, p)
}

func SecretKey(a, B, p *big.Int) (s *big.Int) {
	return new(big.Int).Exp(B, a, p)
}

func NewPair(p *big.Int, g int64) (a, A *big.Int) {
	a = PrivateKey(p)
	return a, PublicKey(a, p, g)
}
