package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"log"
)

const (
	checksumLength = 4
	// version is the hexadecimal representation of the version (currently 0)
	version = byte(0x00)
)

// Wallet struct
type Wallet struct {
	PrivateKey ecdsa.PrivateKeya // Elipitic Curve Digial Signature Algorithm
	PublicKey  []byte
}

func NewKeyPair() (ecdsa.PrivateKey, []byte) {
	curve := elliptic.P256()
	private, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		log.Panic(err)
	}

	pub := append(private.PublicKey.X.Bytes(), private.PublicKey.Y.Bytes()...)
	return *private, pub
}
