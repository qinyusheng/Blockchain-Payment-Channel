package secretbox

import (
	"crypto/rand"
	"io"

	"github.com/btcsuite/btcutil/base58"
	"golang.org/x/crypto/nacl/secretbox"
)

func GenerateSecretKey() string {
	var secretKey [32]byte
	if _, err := io.ReadFull(rand.Reader, secretKey[:]); err != nil {
		panic(err)
	}

	secretKey58 := base58.Encode(secretKey[:])
	return secretKey58
}

func Seal(key string, msg string) string {
	var secretKey [32]byte
	copy(secretKey[:], base58.Decode(key))

	var nonce [24]byte
	if _, err := io.ReadFull(rand.Reader, nonce[:]); err != nil {
		panic(err)
	}

	encrypted := secretbox.Seal(nonce[:], []byte(msg), &nonce, &secretKey)
	encrypted58 := base58.Encode(encrypted)
	return encrypted58
}

func Open(key string, cipher string) (string, bool) {
	var secretKey [32]byte
	copy(secretKey[:], base58.Decode(key))

	var decryptNonce [24]byte
	encrypted := base58.Decode(cipher)
	copy(decryptNonce[:], encrypted[:24])
	decrypted, ok := secretbox.Open(nil, encrypted[24:], &decryptNonce, &secretKey)

	return string(decrypted), ok
}
