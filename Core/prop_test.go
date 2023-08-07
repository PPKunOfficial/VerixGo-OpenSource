package Core_test

import (
	"os"
	"testing"

	"github.com/PPKunOfficial/VerixGo/Core"
)

func TestReadBuildPropFile(t *testing.T) {
	// 测试用的 build.prop 文件路径
	filePath := "build.prop"

	// 创建测试文件
	err := createBuildPropFile(filePath)
	if err != nil {
		t.Fatalf("Failed to create test build.prop file: %v", err)
	}
	defer func() {
		// 删除测试文件
		err := os.Remove(filePath)
		if err != nil {
			t.Fatalf("Failed to remove test build.prop file: %v", err)
		}
	}()

	// 读取 build.prop 文件内容
	propMap, err := Core.ReadBuildPropFile(filePath)
	if err != nil {
		t.Errorf("ReadBuildPropFile() failed: %v", err)
	}

	// 验证读取的键值对数量是否正确
	expectedCount := 3
	if len(propMap) != expectedCount {
		t.Errorf("Read property count is incorrect: expected %d, got %d", expectedCount, len(propMap))
	}

	// 验证读取的键值对是否与预期一致
	expectedProps := map[string]string{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
	}
	for key, expectedValue := range expectedProps {
		value, ok := propMap[key]
		if !ok {
			t.Errorf("Missing property: %s", key)
		} else if value != expectedValue {
			t.Errorf("Property value does not match: expected %s, got %s", expectedValue, value)
		}
	}
}

// 创建测试用的 build.prop 文件
func createBuildPropFile(filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	lines := []string{
		"key1=value1",
		"key2=value2",
		"key3=value3",
		"",
		"# Comment line",
	}

	for _, line := range lines {
		_, err := file.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}

	return nil
}
