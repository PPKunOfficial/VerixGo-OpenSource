package Core

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"os"
)

// SaveData 将结构体数据保存到内存的 byte 变量
func SaveGobToByte(data interface{}) ([]byte, error) {
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	err := encoder.Encode(data)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// LoadData 从内存的 byte 变量加载结构体数据
func LoadGobFromByte(data []byte, target interface{}) error {
	buf := bytes.NewBuffer(data)
	decoder := gob.NewDecoder(buf)
	err := decoder.Decode(target)
	if err != nil {
		return err
	}
	return nil
}

// 保存结构体数据到文件
func SaveGobToFile(filename string, data interface{}) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		return err
	}

	return nil
}

// 从文件加载结构体数据
func LoadGobFromFile(filename string, data interface{}) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := gob.NewDecoder(file)
	err = decoder.Decode(data)
	if err != nil {
		return err
	}

	return nil
}

// 将结构体编码为二进制数据
func EncodeToBytes(data interface{}) ([]byte, error) {
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	err := encoder.Encode(data)
	if err != nil {
		return nil, fmt.Errorf("编码错误: %v", err)
	}
	return buf.Bytes(), nil
}

// 从二进制数据中解码结构体
func DecodeFromBytes(data []byte, result interface{}) error {
	buf := bytes.NewBuffer(data)
	decoder := gob.NewDecoder(buf)
	err := decoder.Decode(result)
	if err != nil {
		return fmt.Errorf("解码错误: %v", err)
	}
	return nil
}
