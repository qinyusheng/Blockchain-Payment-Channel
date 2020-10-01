package envelope

import (
	"digital-envelope/common/box"
	"digital-envelope/common/secretbox"
)

func Seal(msg string, secretKey string, receivePubKey string, senderPriKey string) (string, string) {
	cipher := secretbox.Seal(secretKey, msg)
	encryptedSecretKey := box.Seal(secretKey, receivePubKey, senderPriKey)
	return cipher, encryptedSecretKey
}

func Open(cipher string, encryptedSecretKey string, tempPub string, sessionPri string) (string, bool) {
	decryptedSecretKey, ok := box.Open(encryptedSecretKey, tempPub, sessionPri)
	if !ok {
		return "", ok
	}
	plain, ok := secretbox.Open(decryptedSecretKey, cipher)
	return plain, ok
}
