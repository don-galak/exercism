package diffiehellman

import (
	"crypto/rand"
	"math/big"
)

// Diffie-Hellman-Merkle key exchange
// Private keys should be generated randomly.

var bigTwo = big.NewInt(2)

func PrivateKey(p *big.Int) *big.Int {
	max := new(big.Int).Sub(p, bigTwo)
	random, err := rand.Int(rand.Reader, max)
	if err != nil {
		panic(err)
	}
	return random.Add(bigTwo, random)
}

func PublicKey(private, p *big.Int, g int64) *big.Int {
	return new(big.Int).Exp(big.NewInt(g), private, p)
}

func NewPair(p *big.Int, g int64) (*big.Int, *big.Int) {
	private := PrivateKey(p)
	public := PublicKey(private, p, g)

	return private, public
}

func SecretKey(private1, public2, p *big.Int) *big.Int {
	return new(big.Int).Exp(public2, private1, p)
}
