package pkg

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

func AESEncrypt(src []byte, key []byte, IV []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("key error1", err)
	}
	if len(src) == 0 {
		fmt.Println("plain content empty")
	}
	ecb := cipher.NewCBCEncrypter(block, IV)
	content := (src)
	content = PKCS5Padding(content, block.BlockSize())

	crypted := make([]byte, len(content))
	ecb.CryptBlocks(crypted, content)
	return crypted
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}
