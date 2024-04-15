package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"

)


// func Decrypt(plaintext []byte) string {
// 	bKey := []byte("69aF7&3KY0_kk89@")
// 	bIV := []byte("420#abA%,ZfE79@M")
	
// 	bPlaintext := plaintext
// 	block, _ := aes.NewCipher(bKey)
// 	ciphertext := make([]byte, len(bPlaintext))
// 	mode := cipher.NewCBCDecrypter(block, bIV)
// 	mode.CryptBlocks(ciphertext, bPlaintext)
// 	return hex.EncodeToString(ciphertext)
// }

func Encrypt(plaintext string) string {
	bKey := []byte("69aF7&3KY0_kk89@")
	bIV := []byte("420#abA%,ZfE79@M")
	hexD, _ := hex.DecodeString(plaintext)
	bPlaintext := hexD
	block, _ := aes.NewCipher(bKey)
	ciphertext := make([]byte, len(bPlaintext))
	mode := cipher.NewCBCEncrypter(block, bIV)
	mode.CryptBlocks(ciphertext, bPlaintext)
	return hex.EncodeToString(ciphertext)
}



