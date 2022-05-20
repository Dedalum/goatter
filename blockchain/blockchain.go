package blockchain

import "github.com/dgraph-io/badger"

type BlockChain struct {
	LastHash []byte
	Database *badger.DB
}

const dbPath = "./tmp/blocks"

// NewBlockChain returns a new initialised chain
func NewBlockChain() *BlockChain {
	return &BlockChain{}
}
