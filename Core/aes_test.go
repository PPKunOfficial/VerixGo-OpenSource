package Core_test

import (
	"testing"

	"github.com/PPKunOfficial/VerixGo/Core"
)

func TestAesCBCEncryptAndDecrypt(t *testing.T) {
	// 测试用的明文、密钥和初始化向量
	plaintext := []byte("Hello, World!")
	key := []byte("0123456789abcdef")
	iv := []byte("1234567890acbdef")

	// 加密
	ciphertext, err := Core.AesCBCEncrypt(plaintext, key, iv)
	if err != nil {
		t.Errorf("AesCBCEncrypt() failed: %v", err)
	}

	// 解密
	decrypted, err := Core.AesCBCDecrypt(ciphertext, key, iv)
	if err != nil {
		t.Errorf("AesCBCDecrypt() failed: %v", err)
	}

	// 验证解密结果是否与原始明文一致
	if string(decrypted) != string(plaintext) {
		t.Errorf("Decrypted plaintext does not match original plaintext")
	}
}

func TestPaddingAndUnpaddingPKCS7(t *testing.T) {
	// 测试用的明文和块大小
	plaintext := []byte("Hello, World!")
	blockSize := 16

	// PKCS7 填充
	padded := Core.PaddingPKCS7(plaintext, blockSize)

	// 验证填充后的结果长度是否正确
	if len(padded)%blockSize != 0 {
		t.Errorf("Padding size is incorrect")
	}

	// PKCS7 反填充
	unpadded := Core.UnPaddingPKCS7(padded)

	// 验证反填充后的结果是否与原始明文一致
	if string(unpadded) != string(plaintext) {
		t.Errorf("Unpadded plaintext does not match original plaintext")
	}
}
