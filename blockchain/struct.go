package blockchain

// Block : a single block of data
type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
	Nonce    int
}

// BlockChain : the whole chain. Basically an array of blocks
type BlockChain struct {
	Blocks []*Block
}

// CreateBlock : create the actual block
// in the process, derive it's hash
func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{
		Hash:     []byte{},
		Data:     []byte(data),
		PrevHash: prevHash,
		Nonce: 0,
	}
	proofOfWork := NewProofOfWork(block)
	nonce, hash := proofOfWork.Run()
	block.Hash = hash[:]
	block.Nonce = nonce
	return block
}

// AddBlock : adds a block to the blockchain
// get the last block's hash and use it to create the new block
// then append the new block to the blockchain
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
