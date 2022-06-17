package blockchain

import (
	"fmt"
	"os"
	"runtime"

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

// DBExists checks whether the DB has been initialized
func DBExists(db) bool {
	if _, err := os.Stat(db); os.IsNotExist(err) {
		return false
	}
	return true
}

// NewBlockChain returns a new initialised chain
func InitBlockChain(address string) *BlockChain {

	var lastHash []byte
	if DBExists(dbFile) {
		fmt.Println("blockchain already exists")
		runtime.Goexit()
	}

	opts := badger.DefaultOptions(dbPath)
	db, err := badger.Open(opts)
	Handle(err)

	err = db.Update(func(txn *badger.Txn) error {
		cbtx := CoinbaseTx(address, genesisData)
		genesis := Genesis(cbtx)
		fmt.Println("Genesis created")
		err = txn.Set(genesis.Hash, genesis.Serialize())
		Handle(err)
		err = txn.Set([]byte("lh"), genesis.Hash)

		lastHash = genesis.Hash
		return err
	})

	Handle(err)
	return &BlockChain{lastHash, db}
}

func ContinueBlockChain(address string) *BlockChain {
	if DBexist(dbFile) == flase {
		fmt.Println("No blockchain found, please create one first")
		runtime.Goexit()
	}

	var lastHash []byte
	opts := badger.DefaultOptions(dbPath)
	db, err := badger.Open(opts)
	Handle(err)

	err = db.Update(func(txn *badger.Txn) error {
		iter, err := txn.Get([]byte("lh"))
		Handle(err)
		err = item.Value(func(val []byte) error {
			lastHash = val
			return nil
		})
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
