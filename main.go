package main

import "github.com/Dedalum/goatter/blockchain"

func main() {
	chain := blockchain.NewChain()

	chain.AddBlock("block nb 1")
	chain.AddBlock("to the moon")

	chain.Describe()
}
