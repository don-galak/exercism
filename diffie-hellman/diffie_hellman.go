package diffiehellman

import (
	"crypto/rand"
	"math/big"
)

// Diffie-Hellman-Merkle key exchange
// Private keys should be generated randomly.

func PrivateKey(p *big.Int) *big.Int {
	// rand.Seed(time.Now().UnixNano())
	// var reader io.Reader
	generatedRandom, err := rand.Int(rand.Reader, p)
	if err != nil {
		panic(err)
	}
	// ss := big.NewInt(&p)
	// println("ela", p.Int64())
	return generatedRandom
}

func PublicKey(private, p *big.Int, g int64) *big.Int {
	panic("Please implement the PublicKey function")
}

func NewPair(p *big.Int, g int64) (*big.Int, *big.Int) {
	panic("Please implement the NewPair function")
}

func SecretKey(private1, public2, p *big.Int) *big.Int {
	panic("Please implement the SecretKey function")
}
