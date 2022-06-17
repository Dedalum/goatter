package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"log"
)

type Block struct {
	Hash         []byte
	Transactions []*Transaction
	PrevHash     []byte
	Nonce        int
}

type Chain struct {
	Blocks []*Block
}

// NewChain returns a new initialised chain
func NewChain() *Chain {
	return &Chain{[]*Block{Genesis()}}
}

func (c *Chain) Describe() {
	for _, block := range c.Blocks {
		block.Describe()
	}
}

// NewBlock returns a new block
func NewBlock(txs []*Transaction, prevHash []byte) *Block {
	block := &Block{[]byte{}, txs, prevHash, 0}
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

func Genesis(coinbase *Transaction) *Block {
	return NewBlock([]*Transaction{coinbase}, []byte{})
}

func (b *Block) Serialize() []byte {
	var res bytes.Buffer
	encoder := gob.NewEncoder(&res)

	err := encoder.Encode(b)
	Handle(err)

	return res.Bytes()
}

func Deserialize(data []byte) *Block {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(data))

	err := decoder.Decode(&block)
	Handle(err)
	return &block
}

func Handle(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func (b *Block) Describe() {
	fmt.Printf("previous hash: %x, data: %s, hash: %x\n", b.PrevHash, b.Data, b.Hash)
}

func (b *Block) HashTransactions() []byte {
	var txHashes [][]byte
	var txHash [32]byte

	for _, tx := range b.Transactions {
		txHashes = append(txHashes, tx.ID)
	}
	txHash = sha256.Sum256(bytes.Join(txHashes, []byte))

	return txHash[:]
}
