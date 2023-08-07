package Core

import (
	"bufio"
	"os"
	"strings"
)

func ReadBuildPropFile(filePath string) (map[string]string, error) {
	propMap := make(map[string]string)

	// 打开 build.prop 文件
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// 创建 scanner 以逐行读取文件内容
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// 忽略空行和注释行
		if line == "" || line[0] == '#' {
			continue
		}

		// 解析键值对
		keyValue := strings.SplitN(line, "=", 2)
		if len(keyValue) != 2 {
			// 键值对格式不正确，跳过该行
			continue
		}

		key := strings.TrimSpace(keyValue[0])
		value := strings.TrimSpace(keyValue[1])

		// 将键值对存入 map
		propMap[key] = value
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return propMap, nil
}
