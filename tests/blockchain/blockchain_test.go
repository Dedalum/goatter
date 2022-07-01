package tests

import (
	"os"
	"testing"

	"github.com/Dedalum/goatter/blockchain"
	"github.com/stretchr/testify/assert"
)

func cleanDB() {
	if _, err := os.Stat("./tmp"); !os.IsNotExist(err) {
		os.RemoveAll("./tmp")
	}
}

func TestInitBlockChain(t *testing.T) {
	cleanDB()

	test_address := "test address"
	blockchain := blockchain.InitBlockChain(test_address)
	assert.NotEmpty(t, blockchain.LastHash, "Blockchain should be initialize with a last hash")
}
