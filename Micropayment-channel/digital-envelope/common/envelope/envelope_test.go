package envelope

import (
	"fmt"
	"testing"

	"digital-envelope/common/box"
	"digital-envelope/common/secretbox"
)

// seal: msg, secretKey, sessionPub, ----> cipher, encryptedSecretKey, tempPub
// open: cipher, encryptedSecretKey, tempPub, sessionPri ----> plain, ok
func Test_myEnvelope(t *testing.T) {
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

	cipher, encryptedSecretKey := Seal(msg, secretKey, receiverPubKey, senderPriKey)
	fmt.Println("密文，数字信封：", cipher, encryptedSecretKey)

	plain, ok := Open(cipher, encryptedSecretKey, senderPubKey, receiverPriKey)
	fmt.Println("解密得到：", plain, ok)
}

func Test_envelope(t *testing.T) {
	secretKey := secretbox.GenerateSecretKey()
	fmt.Println(secretKey)

	sessionPub, sessionPri, err := box.GenerateKeyPair()
	fmt.Println(sessionPub, sessionPri, err)

	tempPub, tempPri, err := box.GenerateKeyPair()
	fmt.Println(tempPub, tempPri, err)

	msg := "f*ck envelope"
	fmt.Println(msg)

	cipher := secretbox.Seal(secretKey, msg)
	fmt.Println(cipher)

	encryptedSecretKey := box.Seal(secretKey, sessionPub, tempPri)
	fmt.Println(encryptedSecretKey)

	decryptedSecretKey, ok := box.Open(encryptedSecretKey, tempPub, sessionPri)
	fmt.Println(decryptedSecretKey, ok)

	plain, ok := secretbox.Open(decryptedSecretKey, cipher)
	fmt.Println(plain, ok)
}
