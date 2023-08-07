package Core_test

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"github.com/PPKunOfficial/VerixGo/Core"
	"os"
	"testing"
)

// TestGenerateKey 测试GenerateKey函数
func TestGenerateKey(t *testing.T) {
	// 生成一个临时目录
	dir, err := os.MkdirTemp("", "test")
	if err != nil {
		t.Fatal(err)
	}
	defer func(path string) {
		err := os.RemoveAll(path)
		if err != nil {

		}
	}(dir)

	// 生成密钥对并保存到临时目录
	err = Core.GenerateKey(1024, dir)
	if err != nil {
		t.Fatal(err)
	}

	// 读取私钥文件
	privateKey, err := os.ReadFile(dir + "private.pem")
	if err != nil {
		t.Fatal(err)
	}

	// 读取公钥文件
	publicKey, err := os.ReadFile(dir + "public.pem")
	if err != nil {
		t.Fatal(err)
	}

	// 检查私钥和公钥是否有效
	privateBlock, _ := pem.Decode(privateKey)
	if privateBlock == nil || privateBlock.Type != "RSA Private Key" {
		t.Error("无效的私钥")
	}

	publicBlock, _ := pem.Decode(publicKey)
	if publicBlock == nil || publicBlock.Type != "RSA Public Key" {
		t.Error("无效的公钥")
	}
}

// TestRsaEncrypt 测试RsaEncrypt函数
func TestRsaEncrypt(t *testing.T) {
	// 生成一个临时目录
	dir, err := os.MkdirTemp("", "test")
	if err != nil {
		t.Fatal(err)
	}
	defer func(path string) {
		err := os.RemoveAll(path)
		if err != nil {

		}
	}(dir)

	// 生成密钥对并保存到临时目录
	err = Core.GenerateKey(1024, dir)
	if err != nil {
		t.Fatal(err)
	}

	// 读取公钥文件
	publicKey, err := os.ReadFile(dir + "public.pem")
	if err != nil {
		t.Fatal(err)
	}

	// 定义一个明文数据
	plainText := []byte("Hello, world!")

	// 使用公钥加密明文数据
	cipherText, err := Core.RsaEncrypt(plainText, publicKey)
	if err != nil {
		t.Fatal(err)
	}

	// 检查密文数据是否为空或与明文数据相同
	if len(cipherText) == 0 {
		t.Error("加密失败，密文为空")
	}

	if string(cipherText) == string(plainText) {
		t.Error("加密失败，密文与明文相同")
	}
}

// TestRsaDecrypt 测试RsaDecrypt函数
func TestRsaDecrypt(t *testing.T) {
	// 生成一个临时目录
	dir, err := os.MkdirTemp("", "test")
	if err != nil {
		t.Fatal(err)
	}
	defer func(path string) {
		err := os.RemoveAll(path)
		if err != nil {

		}
	}(dir)

	// 生成密钥对并保存到临时目录
	err = Core.GenerateKey(1024, dir)
	if err != nil {
		t.Fatal(err)
	}

	// 读取私钥文件
	privateKey, err := os.ReadFile(dir + "private.pem")
	if err != nil {
		t.Fatal(err)
	}

	privateBlock, _ := pem.Decode(privateKey)
	if privateBlock == nil || privateBlock.Type != "RSA Private Key" {
		t.Fatal("无效的私钥")
	}

	privateKeyRSA, err := x509.ParsePKCS1PrivateKey(privateBlock.Bytes)
	if err != nil {
		t.Fatal(err)
	}

	publicKeyRSA := privateKeyRSA.PublicKey

	// 定义一个明文数据
	plainText := []byte("Hello, world!")

	// 使用私钥的公钥加密明文数据
	cipherText, err := rsa.EncryptPKCS1v15(rand.Reader, &publicKeyRSA, plainText)
	if err != nil {
		t.Fatal(err)
	}

	// 使用私钥解密密文数据
	newPlainText, err := Core.RsaDecrypt(cipherText, privateKey)
	if err != nil {
		t.Fatal(err)
	}

	// 检查解密后的明文数据是否与原始明文数据相同
	if string(newPlainText) != string(plainText) {
		t.Error("解密失败，明文不匹配")
	}
}
