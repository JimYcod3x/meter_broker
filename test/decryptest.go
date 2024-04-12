package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

func main() {

	ciphertext, err := base64.StdEncoding.DecodeString("9703fef6564f18e815a46a3a9f7b44e6")

	if err != nil {
	}

	fmt.Println(ciphertext)

	// decrypt, err := decrypt([]byte("69aF7&3KY0_kk89@"), "9703fef6564f18e815a46a3a9f7b44e6")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(len(decrypt))
	// fmt.Println(len([]byte(decrypt)))
	// fmt.Println(len([]byte("69aF7&3KY0_kk89@")))
}


