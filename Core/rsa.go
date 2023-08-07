package Core

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"os"
)

// GenerateKey 生成RSA密钥对并保存到指定目录
func GenerateKey(bits int, dirPath string) error {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return err
	}

	privateFile, err := os.Create(dirPath + "private.pem")
	if err != nil {
		return err
	}
	defer func(privateFile *os.File) {
		err := privateFile.Close()
		if err != nil {

		}
	}(privateFile)

	X509PrivateKey := x509.MarshalPKCS1PrivateKey(privateKey)
	privateBlock := pem.Block{Type: "RSA Private Key", Bytes: X509PrivateKey}
	err = pem.Encode(privateFile, &privateBlock)
	if err != nil {
		return err
	}

	publicKey := privateKey.PublicKey
	X509PublicKey, err := x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil {
		return err
	}

	publicFile, err := os.Create(dirPath + "public.pem")
	if err != nil {
		return err
	}
	defer func(publicFile *os.File) {
		err := publicFile.Close()
		if err != nil {

		}
	}(publicFile)

	publicBlock := pem.Block{Type: "RSA Public Key", Bytes: X509PublicKey}
	err = pem.Encode(publicFile, &publicBlock)
	if err != nil {
		return err
	}

	return nil
}

// RsaEncrypt 使用RSA公钥加密数据
func RsaEncrypt(plainText []byte, publicKey []byte) ([]byte, error) {
	block, _ := pem.Decode(publicKey)
	if block == nil || block.Type != "RSA Public Key" {
		return nil, errors.New("无效的公钥")
	}

	publicKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	publicKeyRSA := publicKeyInterface.(*rsa.PublicKey)
	cipherText, err := rsa.EncryptPKCS1v15(rand.Reader, publicKeyRSA, plainText)
	if err != nil {
		return nil, err
	}

	return cipherText, nil
}

// RsaDecrypt 使用RSA私钥解密数据
func RsaDecrypt(cipherText []byte, privateKey []byte) ([]byte, error) {
	block, _ := pem.Decode(privateKey)
	if block == nil || block.Type != "RSA Private Key" {
		return nil, errors.New("无效的私钥")
	}

	privateKeyRSA, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	plainText, err := rsa.DecryptPKCS1v15(rand.Reader, privateKeyRSA, cipherText)
	if err != nil {
		return nil, err
	}

	return plainText, nil
}
