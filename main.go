package main

import (
	"fmt"
	"strconv"

	"github.com/omrfrkazt/blockchain-example/blockchain"
)

func main() {
	chain := blockchain.InitBlockChain()
	chain.AddBlock("first block after Genesis")
	chain.AddBlock("second block after Genesis")
	chain.AddBlock("third block after Genesis")

	for _, block := range chain.Blocks {
		fmt.Printf("previous hash: %x\n ", block.PrevHash)
		fmt.Printf("block data: %s\n", block.Data)
		fmt.Printf("block hash: %x\n ", block.Hash)
		
		pow:= blockchain.NewProof(block)
		fmt.Printf("PoW %s\n",strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}
