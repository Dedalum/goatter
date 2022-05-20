package main

import (
	"fmt"
	"strconv"

	"github.com/Dedalum/goatter/blockchain"
)

func main() {
	chain := blockchain.NewChain()

	chain.AddBlock("block nb 1")
	chain.AddBlock("to the moon")
	chain.AddBlock("and beyond !")

	for _, block := range chain.Blocks {

		fmt.Printf("Previous hash: %x\n", block.PrevHash)
		fmt.Printf("data: %s\n", block.Data)
		fmt.Printf("hash: %x\n", block.Hash)

		pow := blockchain.NewProofOfWork(block)
		fmt.Printf("Pow: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}
