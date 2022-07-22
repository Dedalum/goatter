package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"log"

	"github.com/Dedalum/goatter/blockchain"
	"golang.org/x/crypto/ripemd160"
)

const (
	checksumLength = 4
	// version is the hexadecimal representation of the version (currently 0)
	version = byte(0x00)
)

// Wallet struct
type Wallet struct {
	PrivateKey ecdsa.PrivateKey // Elipitic Curve Digial Signature Algorithm
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

func PublicKeyHash(publicKey []byte) []byte {
	hashedPublicKey := sha256.Sum256(publicKey)

	hasher := ripemd160.New()
	hasher.Write(hashedPublicKey[:]) // never returns an error: https://pkg.go.dev/hash#Hash
	publicRipeMd := hasher.Sum(nil)
	return publicRipeMd
}

func Checksum(ripedMdHash []byte) []byte {
	hash := sha256.Sum256(ripedMdHash)
	hash2 := sha256.Sum256(hash[:])

	return hash2[:checksumLength]
}

func (w *Wallet) Address() []byte {
	pubKeyHash := PublicKeyHash(w.PublicKey)
	versionedHash := append([]byte{version}, pubKeyHash...)
	checksum := Checksum(versionedHash)
	finalHash := append(versionedHash, checksum...)
	address := base58Encode(finalHash)

	return address
}

func NewWallet() *Wallet {
	privKey, pubKey := NewKeyPair()
	return &Wallet{
		privKey,
		pubKey,
	}
}

func (w *Wallet) GetBalance() int {
	address := string(w.Address())
	chain := blockchain.ContinueBlockChain(address)
	defer chain.Database.Close()

	balance := 0
	UTXOs := chain.FindUTXO(address)

	for _, out := range UTXOs {
		balance += out.Value
	}

	return balance
}
