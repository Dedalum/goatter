package blockchain

import (
	"bytes"
	"crypto/sha256"
)

type Chain struct {
	blocks []*Block
}

// NewChain returns a new chain
func NewChain() *Chain {
	return &Chain{}
}

type Block struct {
	Data     []byte
	Hash     []byte
	PrevHash []byte
}

// NewBlock returns a new block
func NewBlock() *Block {
	return &Block{}
}

func (b *Block) DeriveHash() error {
	hash := sha256.Sum256(bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{}))
	b.Hash = hash[:]

	return nil
}
