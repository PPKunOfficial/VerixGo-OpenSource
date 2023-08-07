package main

import (
	"github.com/sirupsen/logrus"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func init() {
	defer func() {
		// 捕获panic时的错误信息
		if err := recover(); err != nil {
			log.Fatalln("Error:", err)
		}
	}()

	// log模块初始化
	initLog()
	log.Infoln("初始日志模块成功")

	initCheckUser()
	log.Infoln("检测用户权限成功")

	// 路径初始化
	initPath()
	log.Infof("初始化路径成功:%v", fPath)

	initVLocal()
	log.Infof("初始化运行时数据成功:%v", vLocal)

	// 设置Gin调试模式
	if !Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	logrus.Infoln("初始化成功")
}

func main() {
	defer func() {
		// 捕获panic时的错误信息
		if err := recover(); err != nil {
			log.Fatalln("Error:", err)
		}
	}()

	// 检查文件是否存在
	_, err := os.Stat(fPath.dbPath)
	if os.IsNotExist(err) {
		log.Printf("文件 %s 不存在", fPath.dbPath)
		vLocal.InstallTime = time.Now().Unix()
		gWriteAuto()
		FirstInstall = true
		log.Info(vLocal)
	} else if err == nil {
		log.Printf("文件 %s 存在", fPath.dbPath)
		m := vLocal.Pmodel
		gReadAuto()
		if vLocal.Pmodel != m {
			FakeExec()
		}
		log.Info(vLocal)
	} else {
		log.Fatalf("检查文件 %s 出错: %v", fPath.dbPath, err)
	}
	go func() {
		defer func() {
			// 捕获panic时的错误信息
			if err := recover(); err != nil {
				log.Fatalln("Error:", err)
			}
		}()
		Joke()
	}()
	if Debug {
		go func() {
			defer func() {
				// 捕获panic时的错误信息
				if err := recover(); err != nil {
					log.Fatalln("Error:", err)
				}
			}()
			DebugHttp()
		}()
	}
	HttpServer()
}
