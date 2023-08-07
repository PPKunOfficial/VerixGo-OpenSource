package Core

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func ReadLinesFromFile(filename string) ([]string, error) {
	var inLines []string
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		inLines = append(inLines, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return inLines, nil
}

// ExecuteShellCommand 执行Shell命令并返回输出结果
func ExecuteShellCommand(command string) (string, error) {
	cmd := exec.Command("sh", "-c", command)

	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("执行命令出错：%v", err)
	}

	outputStr := string(output)
	return outputStr, nil
}

// GetNonSystemAppList 获取非系统应用列表
func GetNonSystemAppList() ([]string, error) {
	command := "pm list packages -3 | sed -e 's/^package://'"
	output, err := ExecuteShellCommand(command)
	if err != nil {
		return nil, err
	}
	// 根据回车切割字符串
	return strings.Split(output, "\n"), nil
}
