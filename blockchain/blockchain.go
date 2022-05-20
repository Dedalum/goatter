package blockchain

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

type Chain struct {
	Blocks []*Block
}

// NewChain returns a new initialised chain
func NewChain() *Chain {
	return &Chain{[]*Block{Genesis()}}
}

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
	Nonce    int
}

// NewBlock returns a new block
func NewBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash, 0}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

func (c *Chain) AddBlock(data string) {
	prevHash := c.Blocks[len(c.Blocks)-1].Hash
	c.Blocks = append(c.Blocks, NewBlock(data, prevHash))
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
	for _, block := range c.Blocks {
		block.Describe()
	}
}
