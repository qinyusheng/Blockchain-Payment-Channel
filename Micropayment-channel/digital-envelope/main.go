package main

import (
	"fmt"
	"digital-envelope/common/box"
	"digital-envelope/common/envelope"
	"digital-envelope/common/secretbox"
	"os"
)

func main() {
	secretKey := secretbox.GenerateSecretKey() // 对称密钥，用于加密明文消息
	fmt.Println("对称密钥：", secretKey)

	senderPubKey, senderPriKey, err := box.GenerateKeyPair() // 发送方 公私钥
	if err != nil {
		panic(err)
	}
	fmt.Println("发送方公私钥：", senderPubKey, senderPriKey)

	receiverPubKey, receiverPriKey, err := box.GenerateKeyPair() // 接收方 公私钥
	if err != nil {
		panic(err)
	}
	fmt.Println("接收方公私钥：", receiverPubKey, receiverPriKey)

	msg := "f*ck envelope seal and open ?" // 明文消息
	fmt.Println("明文消息：", msg)

	cipher, encryptedSecretKey := envelope.Seal(msg, secretKey, receiverPubKey, senderPriKey)
	fmt.Println("密文，数字信封：", cipher, encryptedSecretKey)

	plain, ok := envelope.Open(cipher, encryptedSecretKey, senderPubKey, receiverPriKey)
	fmt.Println("解密得到：", plain, ok)
}
