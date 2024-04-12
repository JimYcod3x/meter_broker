package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
)

func main() {

	// ciphertext, err := base64.StdEncoding.DecodeString("9703fef6564f18e815a46a3a9f7b44e6")

	// if err != nil {
	// }

	// fmt.Println(ciphertext)

	bytes := []byte("\x97\x03\xfe\xf6VO\x18\xe8\x15\xa4j:\x9f{D\xe6")
	hexString := hex.EncodeToString(bytes)
	fmt.Println(hexString)
	fmt.Println(Ase256("\x97\x03\xfe\xf6VO\x18\xe8\x15\xa4j:\x9f{D\xe6", "69aF7&3KY0_kk89@", "420#abA%,ZfE79@M", 128))
	fmt.Println(AseE256("6a004a23500000780000000000000000", "69aF7&3KY0_kk89@", "420#abA%,ZfE79@M", 128))

	

	// decrypt, err := decrypt([]byte("69aF7&3KY0_kk89@"), "9703fef6564f18e815a46a3a9f7b44e6")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(len(decrypt))
	// fmt.Println(len([]byte(decrypt)))
	// fmt.Println(len([]byte("69aF7&3KY0_kk89@")))
}

func Ase256(plaintext string, key string, iv string, blockSize int) string {
	bKey := []byte(key)
	bIV := []byte(iv)
	
	bPlaintext := []byte(plaintext)
	block, _ := aes.NewCipher(bKey)
	ciphertext := make([]byte, len(bPlaintext))
	mode := cipher.NewCBCDecrypter(block, bIV)
	mode.CryptBlocks(ciphertext, bPlaintext)
	return hex.EncodeToString(ciphertext)
}

func AseE256(plaintext string, key string, iv string, blockSize int) string {
	bKey := []byte(key)
	bIV := []byte(iv)
	hexD, _ := hex.DecodeString(plaintext)
	bPlaintext := hexD
	block, _ := aes.NewCipher(bKey)
	ciphertext := make([]byte, len(bPlaintext))
	mode := cipher.NewCBCEncrypter(block, bIV)
	mode.CryptBlocks(ciphertext, bPlaintext)
	return hex.EncodeToString(ciphertext)
}



