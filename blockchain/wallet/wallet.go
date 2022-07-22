//wallet.go
package wallet

import (
    "crypto/ecdsa"
    "crypto/elliptic"
    "crypto/rand"
    "crypto/sha256"
    "log"
    "github.com/Dedalum/goatter/blockchain/wallet"
    "golang.org/x/crypto/ripemd160"
)

const (
    checsumLength = 4
    // hexadecimal representation od 0
    version = byte(0x00)
)

type Wallet struct {
    // ecdsa = eliptical curve digital signature algorithm
    PrivateKey ecdsa.PrivateKey
    PublicKey []byte
}

func NewKeyPair() (ecdsa.PrivateKey, []byte) {
    curve := elliptic.P256()

    priate, err := ecdsa.GenerateKey(curve, rand.Reader)
    if err != nil {
        log.Panic(err)
    }

    pub := append(private.PublicKey.X.Bytes(), private.PublicKey.Y.Bytes()...)

    return *private, pub
}

func PublicKeyHash(publickey []byte) []byte {
    hashedPublicKey := sha256.Sum256(publickey)
    hasher := ripemd160.New()
    _, err := hasher.Write(hashedPublicKey[:])
    if err != nil {
        log.Panic(err)
    }
    publicRipeMd := hasher.Sum(nil)
    return publicRipeMd
}

func CheckSum(ripeMdHash []byte) []byte {
    firstHash := sha256.Sum256(ripeMdHash)
    secondHash := sha256.Sum256(firstHash[:])
    return secondHash[:checsumLength]
}

func (w *Wallet) Address() []byte {
    // step 1 and 2
    pubHash := PublicKeyHash(w.PublicKey)
    // step 3
    versionedHash := append([]byte{version}, pubHash...)
    // step 4
    checksum := CheckSum(versionedHash)
    // step 5
    finalHash := append(versionedHash, checksum...)
    // step 6
    address := base58Encode(finalHash)
    return address
}

func MakeWallet() *Wallet {
    privateKey, publicKey := NewKeyPair()
    wallet := Wallet{privateKey, publicKey}
    return &wallet
}
