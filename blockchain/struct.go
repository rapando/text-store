package blockchain

import (
	"bytes"
	"crypto/sha256"
)

// Block : a single block of data
type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

// BlockChain : the whole chain. Basically an array of blocks
type BlockChain struct {
	Blocks []*Block
}

// DeriveHash : generates the hash of the current block
// It joins the previous block's info with the current block's data
func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte {b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

// CreateBlock : create the actual block
// in the process, derive it's hash
func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{
		Hash: []byte{},
		Data: []byte(data),
		PrevHash: prevHash,
	}
	block.DeriveHash()
	return block
}

// AddBlock : adds a block to the block chain
// get the last block's hash and use it to create the new block
// then append the new block to the block chain
func (c *BlockChain) AddBlock(data string) {
	prevBlock := c.Blocks[len(c.Blocks)-1]
	newBlock := CreateBlock(data, prevBlock.Hash)
	c.Blocks = append(c.Blocks, newBlock)
}

// Genesis : create our 'genesis' block
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

// InitBlockChain : create a new block chain with the genesis block
func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}
