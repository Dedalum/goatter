package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
)

const reward = 100

type Transaction struct {
	Id      []byte
	Inputs  []TxInput
	Outputs []TxOutput
}

func CoinbaseTx(toAddress, data string) *Transaction {
	if data == "" {
		data = fmt.Sprintf("Coins to %s", toAddress)
	}
	//Since this is the "first" transaction of the block, it has no previous output to reference.
	//This means that we initialize it with no ID, and it's OutputIndex is -1
	txIn := TxInput{[]byte{}, -1, data}

	txOut := TxOutput{reward, toAddress}

	tx := Transaction{nil, []TxInput{txIn}, []TxOutput{txOut}}

	return &tx
}

func (tx *Transaction) SetID() {
	var encoded bytes.Buffer
	var hash [32]byte

	encoder := gob.NewEncoder(&encoded)
	err := encoder.Encode(tx)
	Handle(err)

	hash = sha256.Sum256(encoded.Bytes())
	tx.Id = hash[:]

}

func (tx *Transaction) IsCoinbase() bool {
	//This checks a transaction and will only return true if it is a newly minted "coin"
	return len(tx.Inputs) == 1 && len(tx.Inputs[0].ID) == 0 && tx.Inputs[0].Out == -1
}

func NewTransaction(from, to string, amount int, chain *BlockChain) *Transaction {
	var inputs []TxInput
	var outputs []TxOutput

	//STEP 1
	acc, validOutputs := chain.FindSpendableOutputs(from, amount)

	//STEP 2
	if acc < amount {
		log.Panic("Error: Not enough funds!")
	}

	//STEP 3
	for txid, outs := range validOutputs {
		txID, err := hex.DecodeString(txid)
		Handle(err)

		for _, out := range outs {
			input := TxInput{txID, out, from}
			inputs = append(inputs, input)
		}
	}

	outputs = append(outputs, TxOutput{amount, to})

	//STEP 4
	if acc > amount {
		outputs = append(outputs, TxOutput{acc - amount, from})
	}

	//STEP 5
	tx := Transaction{nil, inputs, outputs}
	//STEP 6
	tx.SetID()

	return &tx
}

// Serialise the transaction to send over network
func (tx *Transaction) Serialise() []byte {
	serialised_tx, err := json.Marshal(tx)
	if err != nil {
		log.Panic() //TODO: use logger
	}

	return serialised_tx
}

func Unserialise(transaction []byte) *Transaction {
	tx := &Transaction{}
	if err := json.Unmarshal(transaction, tx); err != nil {
		log.Panic()
	}
	return tx
}
