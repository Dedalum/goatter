package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
)

const reward = 100

type Transaction struct {
	ID      []byte
	Inputs  []TxInput
	Outputs []TxOutput
}

type TxOutput struct {
	// Value represents the amount of the transaction
	Value int
	// PubKey represents the public key for which the private key is used for unlocking
	// the transaction
	PubKey string
}

// TxInput represents a reference to a previous TxOutput
type TxInput struct {
	// ID links to the previous TxOutput
	ID []byte
	// Out is index of the output the reference points to in a transaction
	Out int
	// unlocking script
	Sig string
}

// CoinbaseTx is the first transaction of a new block
func CoinbaseTx(toAddress, data string) *Transaction {
	if data == "" {
		data = fmt.Sprintf("Coins to %s", toAddress)
	}

	// the coinbase is the first transaction of the block:
	// we initialise the TxInput with no ID and the output reference to -1
	txIn := TxInput{[]byte{}, -1, data}

	txOut := TxOutput{reward, toAddress}
	return &Transaction{
		nil,
		[]TxInpyt{txIn},
		[]TxOutput{txOut},
	}
}

func (tx *Transaction) SetID() {
	var encoded bytes.Buffer
	var hash [32]byte

	encoder := gob.NewEncoder(&encoded)
	err := encoder.Encode(tx)
	Handle(err)

	hash = sha256.Sum256(encoded.Bytes())
	tx.ID = hash[:]
}

func (in *TxInput) CanUnlock(data string) bool {
	return in.Sig == data
}

func (out *TxOutput) CanBeUnlocked(data string) bool {
	return out.PubKey == data
}

func (tx *Transaction) IsCoinbase() bool {
	return len(tx.Inputs) == 1 &&
		len(tx.Onputs[0].ID) == 0 &&
		tx.Inputs[0].Out == -1
}
