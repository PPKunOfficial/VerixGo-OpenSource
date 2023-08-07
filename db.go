package main

import (
	"io"
	"os"

	"github.com/PPKunOfficial/VerixGo/Core"
)

func gWriteAuto() {
	// 创建或打开 gob 文件
	file, err := os.OpenFile(fPath.dbPath, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Error(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	b, err := Core.SaveGobToByte(vLocal)
	if err != nil {
		log.Fatalf("保存gob出现异常:%s", err)
	}
	b, err = Core.AesCBCEncrypt(b, []byte(aesPassword), []byte(iv))
	if err != nil {
		log.Fatalf("加密出现异常:%s", err)
	}
	_, err = file.Write(b)
	if err != nil {
		log.Fatalf("写入gob出现异常:%s", err)
	}
}
func gReadAuto() {

	// 创建或打开 gob 文件
	file, err := os.Open(fPath.dbPath)
	if err != nil {
		log.Error(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	b, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("读取gob异常:%s", err)
	}
	b, err = Core.AesCBCDecrypt(b, []byte(aesPassword), []byte(iv))
	if err != nil {
		log.Fatalf("解密出现异常:%s", err)
	}
	err = Core.DecodeFromBytes(b, &vLocal)
	if err != nil {
		log.Fatalf("解码gob出现异常:%s", err)
	}
}
