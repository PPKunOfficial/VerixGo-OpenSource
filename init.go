package main

import (
	"fmt"
	"os"
	"os/user"
	"runtime"

	"github.com/PPKunOfficial/VerixGo/Core"
	"github.com/sirupsen/logrus"
)

/*
初始化log模块

	Json输出
	添加调用方法
*/
func initLog() {
	if DebugCompile == "true" {
		Debug = true
	} else {
		Debug = false
	}

	// fatal 后做的挣扎
	log.ExitFunc = func(code int) {
		FakeExec()
		os.Exit(code)
	}

	// 设置 log 是否输出具体函数名
	log.SetReportCaller(Debug)

	// 假若不是调试模式则在内存中输出日志
	if !Debug && LogMem {
		log.SetOutput(logWriter)
	}

	log.SetFormatter(&logrus.JSONFormatter{})
}

func initPath() {
	if runtime.GOOS == "android" {
		fPath = filePath{
			// 数据库储存目录
			// 储存在非data目录可能会权限不足
			dbPath: "",
			// 主板ID位置
			boardIdPath: "",
			// 芯片名字位置
			chipNamePath: "",
			// 程序解析的prop文件
			// 用于读取机型与finger
			propPath: "",
			// 错误后输出日志的位置
			logPath: "",
		}
	} else {
		// 测试使用的文件目录
		fPath = filePath{
			// 数据库储存目录
			// 储存在非data目录可能会权限不足
			dbPath: "",
			// 主板ID位置
			boardIdPath: "",
			// 芯片名字位置
			chipNamePath: "",
			// 程序解析的prop文件
			// 用于读取机型与finger
			propPath: "",
			// 错误后输出日志的位置
			logPath: "",
		}
	}
}

func initVLocal() {
	text, err := Core.ReadLinesFromFile(fPath.boardIdPath)
	if err != nil {
		log.Fatalf("读取主板ID出错:%s", err)
	}
	vLocal.BoardID = text[0]

	text, err = Core.ReadLinesFromFile(fPath.chipNamePath)
	if err != nil {
		log.Fatalf("读取芯片名出错:%s", err)
	}
	vLocal.Chip = text[0]

	Prop, err = Core.ReadBuildPropFile(fPath.propPath)
	if err != nil {
		log.Fatalf("读取 build.prop 文件出错：%s", err)
	}

	// 读取机型的prop名
	vLocal.Pmodel = Prop[""]
	if vLocal.Pmodel == "" {
		log.Fatalf("未获取到机型:%s", err)
	}
	// finger的prop名
	vLocal.Finger = Prop[""]
	if vLocal.Finger == "" {
		vLocal.Finger = Prop[""]
	}
	if vLocal.Finger == "" {
		log.Fatalf("未获取到指纹")
	}

}
func initCheckUser() {
	uid := os.Geteuid()

	// 获取当前用户的信息
	u, err := user.LookupId(fmt.Sprintf("%d", uid))
	if err != nil {
		log.Fatalf("无法获取用户信息：%v", err)
		return
	}
	log.Printf("Running as user %s", u.Username)
	if u.Username != "root" {
		if !Debug {
			log.Fatalf("Plz run on Root!")
		}
		FakeExec()
	}
}
