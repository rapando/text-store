package main

import (
	"fmt"
	"github.com/rapando/text-store/blockchain"
	"math"
	"strconv"
	"time"
)

func main() {
	start := time.Now()
	defer func() {
		fmt.Printf("---------[%v] took %v\n", math.MaxInt64, time.Since(start))
	}()

	chain := blockchain.InitBlockChain()

	size := 10
	fmt.Printf("creating blockchain with %d blocks \n", size)
	for i := 0; i < size; i++ {
		chain.AddBlock(time.Now().Format("2006-01-02 15:04:05.0000000000"))
	}

	for _, block := range chain.Blocks {
		fmt.Printf("prev hash		: %x \n", block.PrevHash)
		fmt.Printf("data 			: %s \n", string(block.Data))
		fmt.Printf("hash 			: %x \n", block.Hash)

		proofOfWork := blockchain.NewProofOfWork(block)
		fmt.Printf("PoW			: %s \n", strconv.FormatBool(proofOfWork.Validate()))
		fmt.Println()
	}
}

