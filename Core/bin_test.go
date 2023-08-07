package Core_test

import (
	"os"
	"testing"

	"github.com/PPKunOfficial/VerixGo/Core"
)

type TestData struct {
	Name  string
	Age   int
	Email string
}

func TestSaveAndLoadGobToByte(t *testing.T) {
	// 测试用的数据
	data := TestData{
		Name:  "John Doe",
		Age:   30,
		Email: "johndoe@example.com",
	}

	// 保存数据到 byte 变量
	bytes, err := Core.SaveGobToByte(data)
	if err != nil {
		t.Errorf("SaveGobToByte() failed: %v", err)
	}

	// 从 byte 变量加载数据
	var loadedData TestData
	err = Core.LoadGobFromByte(bytes, &loadedData)
	if err != nil {
		t.Errorf("LoadGobFromByte() failed: %v", err)
	}

	// 验证加载的数据是否与原始数据一致
	if loadedData != data {
		t.Errorf("Loaded data does not match original data")
	}
}

func TestSaveAndLoadGobToFile(t *testing.T) {
	// 测试用的数据
	data := TestData{
		Name:  "John Doe",
		Age:   30,
		Email: "johndoe@example.com",
	}

	// 保存数据到文件
	filename := "test_data.gob"
	err := Core.SaveGobToFile(filename, data)
	if err != nil {
		t.Errorf("SaveGobToFile() failed: %v", err)
	}
	defer func() {
		// 删除测试文件
		err := os.Remove(filename)
		if err != nil {
			t.Errorf("Failed to remove test file: %v", err)
		}
	}()

	// 从文件加载数据
	var loadedData TestData
	err = Core.LoadGobFromFile(filename, &loadedData)
	if err != nil {
		t.Errorf("LoadGobFromFile() failed: %v", err)
	}

	// 验证加载的数据是否与原始数据一致
	if loadedData != data {
		t.Errorf("Loaded data does not match original data")
	}
}

func TestEncodeAndDecodeFromBytes(t *testing.T) {
	// 测试用的数据
	data := TestData{
		Name:  "John Doe",
		Age:   30,
		Email: "johndoe@example.com",
	}

	// 编码数据为二进制
	bytes, err := Core.EncodeToBytes(data)
	if err != nil {
		t.Errorf("EncodeToBytes() failed: %v", err)
	}

	// 解码二进制数据
	var decodedData TestData
	err = Core.DecodeFromBytes(bytes, &decodedData)
	if err != nil {
		t.Errorf("DecodeFromBytes() failed: %v", err)
	}

	// 验证解码的数据是否与原始数据一致
	if decodedData != data {
		t.Errorf("Decoded data does not match original data")
	}
}
