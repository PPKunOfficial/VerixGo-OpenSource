package Core_test

import (
	"os"
	"runtime"
	"testing"

	"github.com/PPKunOfficial/VerixGo/Core"
)

func TestReadLinesFromFile(t *testing.T) {
	// 测试用的文件名
	filename := "test.txt"

	// 创建测试文件
	err := createTestFile(filename)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	defer func() {
		// 删除测试文件
		err := os.Remove(filename)
		if err != nil {
			t.Fatalf("Failed to remove test file: %v", err)
		}
	}()

	// 读取文件内容
	lines, err := Core.ReadLinesFromFile(filename)
	if err != nil {
		t.Errorf("ReadLinesFromFile() failed: %v", err)
	}

	// 验证读取的行数是否正确
	expectedLines := []string{"Line 1", "Line 2", "Line 3"}
	if len(lines) != len(expectedLines) {
		t.Errorf("Read lines count is incorrect: expected %d, got %d", len(expectedLines), len(lines))
	}

	// 验证读取的每一行是否与预期一致
	for i, line := range lines {
		if line != expectedLines[i] {
			t.Errorf("Read line does not match: expected %s, got %s", expectedLines[i], line)
		}
	}
}

func TestExecuteShellCommand(t *testing.T) {
	if runtime.GOOS != "windows" {
		// 执行 Shell 命令
		output, err := Core.ExecuteShellCommand("echo Hello, World!")
		if err != nil {
			t.Errorf("ExecuteShellCommand() failed: %v", err)
		}

		// 验证输出结果是否符合预期
		expectedOutput := "Hello, World!\n"
		if output != expectedOutput {
			t.Errorf("Output does not match: expected %s, got %s", expectedOutput, output)
		}
	}
}

// 创建测试文件
func createTestFile(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	lines := []string{"Line 1", "Line 2", "Line 3"}
	for _, line := range lines {
		_, err = file.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}

	return nil
}
