package blockchain

import (
	"fmt"
	"os"

	"github.com/dgraph-io/badger"
)

type BlockChain struct {
	LastHash []byte
	Database *badger.DB
}

type BlockChainIterator struct {
	CurrentHash []byte
	Database    *badger.DB
}

const (
	dbPath       = "./tmp/blocks"
	dbFile       = "./tmp/blocks/MANIFEST"                  // this filie can be used to verify the blokchain exists
	genesisDatsa = "Goatter first transaction from Genesis" // arbitrary data
)

// NewBlockChain returns a new initialised chain
func InitBlockChain() *BlockChain {

	// 1
	var lastHash []byte
	opts := badger.DefaultOptions(dbPath)

	db, err := badger.Open(opts)
	Handle(err)

	// 2
	err = db.Update(func(txn *badger.Txn) error {
		// lh = last hash
		item, err := txn.Get([]byte("lh"))
		if err == badger.ErrKeyNotFound {
			fmt.Println("No existing blockchain found")
			genesis := Genesis()
			fmt.Println("Genesis proved")
			err = txn.Set(genesis.Hash, genesis.Serialize())
			Handle(err)
			err = txn.Set([]byte("lh"), genesis.Hash)

			lastHash = genesis.Hash
			return err
		} else {
			// 3
			Handle(err)
			err = item.Value(func(val []byte) error {
				lastHash = val
				return nil
			})
			Handle(err)
			return err
		}
	})
	Handle(err)

	return &BlockChain{lastHash, db}
}

func (chain *BlockChain) AddBlock(data string) {
	var lastHash []byte
	err := chain.Database.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte("lh"))
		Handle(err)
		err = item.Value(func(val []byte) error {
			lastHash = val
			return nil
		})
		Handle(err)
		return err
	})
	Handle(err)

	newBlock := NewBlock(data, lastHash)
	err = chain.Database.Update(func(transaction *badger.Txn) error {
		err := transaction.Set(newBlock.Hash, newBlock.Serialize())
		Handle(err)
		err = transaction.Set([]byte("lh"), newBlock.Hash)
		chain.LastHash = newBlock.Hash
		return err
	})
	Handle(err)
}

func (chain *BlockChain) Iterator() *BlockChainIterator {
	iterator := BlockChainIterator{chain.LastHash, chain.Database}
	return &iterator
}

func (iterator *BlockChainIterator) Next() *Block {
	var block *Block

	err := iterator.Database.View(func(txn *badger.Txn) error {
		item, err := txn.Get(iterator.CurrentHash)
		Handle(err)

		err = item.Value(func(val []byte) error {
			block = Deserialize(val)
			return nil
		})
		return err
	})
	Handle(err)
	iterator.CurrentHash = block.PrevHash

	return block
}

// DBExists checks whether a DB has been initialized or not
func DBExists(db) bool {
	if _, err := os.Stat(db); os.IsNotExist(err) {
		return false
	}
	return true
}
