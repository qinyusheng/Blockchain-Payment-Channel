package main

import (
	"digital-envelope/common/box"
	"digital-envelope/common/secretbox"
)


//参数列表：（明文，对称密钥，接收方公钥，发送方密钥）
//返回值：密文，数字信封
func Seal(msg string, secretKey string, receivePubKey string, senderPriKey string) (string, string) {
	cipher := secretbox.Seal(secretKey, msg)
	encryptedSecretKey := box.Seal(secretKey, receivePubKey, senderPriKey)
	return cipher, encryptedSecretKey
}
//参数列表：（密文，数字信封，发送方公钥，接收方私钥）
//返回值：明文，是否解密成功
func Open(cipher string, encryptedSecretKey string, senderPubKey string, receiverPriKey string) (string, bool) {
	decryptedSecretKey, ok := box.Open(encryptedSecretKey, senderPubKey, receiverPriKey)
	if !ok {
		return "", ok
	}
	plain, ok := secretbox.Open(decryptedSecretKey, cipher)
	return plain, ok
}
