package Core

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
)

/*
AES-128：密钥长度为 128 位（16 字节）
AES-192：密钥长度为 192 位（24 字节）
AES-256：密钥长度为 256 位（32 字节）
*/
func GenerateAESKey(keySize int) ([]byte, error) {
	key := make([]byte, keySize)
	_, err := rand.Read(key)
	if err != nil {
		return nil, err
	}

	return key, nil
}

// aesCBCEncrypt AES/CBC/PKCS7Padding 加密
func AesCBCEncrypt(plaintext []byte, key []byte, iv []byte) ([]byte, error) {
	// AES
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Printf("初始化AES失败:%s", err)
	}

	// PKCS7 填充
	plaintext = PaddingPKCS7(plaintext, aes.BlockSize)

	// CBC 加密
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(plaintext, plaintext)

	return plaintext, nil
}

// aesCBCDecrypt AES/CBC/PKCS7Padding 解密
func AesCBCDecrypt(ciphertext []byte, key []byte, iv []byte) ([]byte, error) {
	// AES
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Printf("初始化AES失败:%s", err)
	}

	if len(ciphertext)%aes.BlockSize != 0 {
		fmt.Println("ciphertext is not a multiple of the block size")
	}

	// CBC 解密
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)

	// PKCS7 反填充
	result := UnPaddingPKCS7(ciphertext)
	return result, nil
}

// PKCS7 填充
func PaddingPKCS7(plaintext []byte, blockSize int) []byte {
	paddingSize := blockSize - len(plaintext)%blockSize
	paddingText := bytes.Repeat([]byte{byte(paddingSize)}, paddingSize)
	return append(plaintext, paddingText...)
}

// PKCS7 反填充
func UnPaddingPKCS7(s []byte) []byte {
	length := len(s)
	if length == 0 {
		return s
	}
	unPadding := int(s[length-1])
	return s[:(length - unPadding)]
}
