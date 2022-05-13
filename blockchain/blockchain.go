package blockchain

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

type Chain struct {
	blocks []*Block
}

// NewChain returns a new initialised chain
func NewChain() *Chain {
	return &Chain{[]*Block{Genesis()}}
}

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

// NewBlock returns a new block
func NewBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash}
	block.DeriveHash()

	return block
}

func (c *Chain) AddBlock(data string) {
	prevHash := c.blocks[len(c.blocks)-1].Hash
	c.blocks = append(c.blocks, NewBlock(data, prevHash))
}

func Genesis() *Block {
	return NewBlock("GoatterGenesis", []byte{})
}

func (b *Block) DeriveHash() {
	hash := sha256.Sum256(bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{}))
	b.Hash = hash[:]
}

func (b *Block) Describe() {
	fmt.Printf("previous hash: %x, data: %s, hash: %x\n", b.PrevHash, b.Data, b.Hash)
}

func (c *Chain) Describe() {
	for _, block := range c.blocks {
		block.Describe()
	}
}
