package box

import (
	cryptoRand "crypto/rand"
	"io"

	"github.com/btcsuite/btcutil/base58"
	"golang.org/x/crypto/nacl/box"
)

func GenerateKeyPair() (string, string, error) {
	senderPublicKey, senderPrivateKey, err := box.GenerateKey(cryptoRand.Reader)
	if err != nil {
		return "", "", err
	}

	senderPublicKey58 := base58.Encode(senderPublicKey[:])
	senderPrivateKey58 := base58.Encode(senderPrivateKey[:])
	return senderPublicKey58, senderPrivateKey58, nil
}

func Seal(msg string, recipientPublicKey string, senderPrivateKey string) string {
	recipientPublicKey58d := base58.Decode(recipientPublicKey)
	senderPrivateKey58d := base58.Decode(senderPrivateKey)
	recipientPub := new([32]byte)
	senderPri := new([32]byte)
	copy(recipientPub[:], recipientPublicKey58d)
	copy(senderPri[:], senderPrivateKey58d)

	var nonce [24]byte
	if _, err := io.ReadFull(cryptoRand.Reader, nonce[:]); err != nil {
		panic(err)
	}

	// This encrypts msg and appends the result to the nonce.
	encrypted := box.Seal(nonce[:], []byte(msg), &nonce, recipientPub, senderPri)
	encrypted58 := base58.Encode(encrypted)
	return encrypted58
}

func Open(cipher string, senderPublicKey string, recipientPrivateKey string) (string, bool) {
	senderPublicKey58d := base58.Decode(senderPublicKey)
	recipientPrivateKey58d := base58.Decode(recipientPrivateKey)
	senderPub := new([32]byte)
	recipientPri := new([32]byte)
	copy(senderPub[:], senderPublicKey58d)
	copy(recipientPri[:], recipientPrivateKey58d)

	encrypted := base58.Decode(cipher)
	var decryptNonce [24]byte
	copy(decryptNonce[:], encrypted[:24])
	decrypted, ok := box.Open(nil, encrypted[24:], &decryptNonce, senderPub, recipientPri)

	return string(decrypted), ok
}
