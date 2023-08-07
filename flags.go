package main

import (
	"bytes"
	"github.com/sirupsen/logrus"
)

var (
	// Debug 调试模式
	// 关闭后将不会显示报错
	Debug        = false
	DebugCompile string
	LogMem       = true

	// 个性化变量
	/*
		面板
			panel
			Panel结构体
				n 验证后缀
				name 面板名字
				port 面板端口

			Aes/iv加解密密码

			wait 等待时间
	*/
	panel = Panel{
		// 云端文件后缀
		n: "",
		// 验证面板名字
		name: "",
		// 面板端口
		port: "",
	}

	// AES密码与IV 采用aes-cbc-256
	// 密钥与IV长度自己填符合aes-cbc-256长度的
	aesPassword = ""
	iv          = ""
	wait        = 3600

	// 运行时变量 不可修改
	logByteBuffer bytes.Buffer
	logWriter     = &ByteWriter{buffer: &logByteBuffer}
	log           = logrus.New()
	vLocal        MPhone
	FirstInstall  bool
	fPath         filePath
	Prop          map[string]string
)

type Panel struct {
	n    string
	name string
	port string
}

// MPhone 手机核心信息
type MPhone struct {
	BoardID     string
	Pmodel      string
	Chip        string
	Finger      string
	Verify      bool
	CheckTime   int64
	InstallTime int64
}

// 校验文件路径
type filePath struct {
	dbPath       string
	boardIdPath  string
	chipNamePath string
	propPath     string
	logPath      string
}
