package blockchain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeriveHash(t *testing.T) {

	expectedHash := []byte("1234")
	b := Block{
		PrevHash: []byte{},
		Data:     []byte("test"),
	}
	err := b.DeriveHash()
	assert.Nil(t, err)
	assert.Equal(b.Hash)
}
