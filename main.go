package main

import (
	"fmt"
	"github.com/rapando/text-store/blockchain"
)

func main() {

	chain := blockchain.InitBlockChain()

	chain.AddBlock ("first block")
	chain.AddBlock ("second block")
	chain.AddBlock ("third block")

	for _, block := range chain.Blocks {
		fmt.Printf("prev hash		: %x \n", block.PrevHash)
		fmt.Printf("data 			: %s \n", string(block.Data))
		fmt.Printf("hash 			: %x \n", block.Hash)
		fmt.Println()
	}
}
