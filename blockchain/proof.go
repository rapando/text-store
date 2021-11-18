/*
1. Grab data from the block
2. Create a nonce
3. Create a hash of the data from the block + nonce
4. Check the proceeding hash to see if it meets the requirements listed above
*/

package blockchain

import (
	"bytes"
	"crypto/sha256"
	"math"
	"math/big"
)

const Difficulty = 1

type ProofOfWork struct {
	Block  *Block
	Target *big.Int
}

// NewProofOfWork : create a new proof of work for a block
func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-Difficulty))
	return &ProofOfWork{b, target}
}

// InitNonce : create a nonce
func (p *ProofOfWork) InitNonce(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			p.Block.PrevHash,
			p.Block.Data,
			ToHex(int64(nonce)),
			ToHex(int64(Difficulty)),
		},
		[]byte{},
	)
	return data
}

// Run : runs the proof of work algorithm for the block.
func (p *ProofOfWork) Run() (int, []byte) {
	var intHash big.Int
	var hash [32]byte

	// create a loop of how many times it is from 0 to maxInt64
	nonce := 0
	for nonce = 0; nonce < math.MaxInt64; nonce++ {
		data := p.InitNonce(nonce)
		hash := sha256.Sum256(data)

		//fmt.Printf("\r [%d] %x", nonce, hash)
		intHash.SetBytes(hash[:])

		if intHash.Cmp(p.Target) == -1 {
			break
		}
	}
	//fmt.Println()
	return nonce, hash[:]
}


// Validate : validate our own proof of work
func (p *ProofOfWork) Validate() bool {
	var intHash big.Int

	data := p.InitNonce(p.Block.Nonce)

	hash := sha256.Sum256(data)
	intHash.SetBytes(hash[:])

	return intHash.Cmp(p.Target) == -1
}