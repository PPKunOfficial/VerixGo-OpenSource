# VerixGo-OpenSource 构建指南
## A. 初始准备
### 1. 编译环境
将代码仓库完整克隆到本地，然后将`Android NDK`工具链添加到环境变量

根据Go官网及网上教程安装并配置好Go语言编译环境(注意：配置`go mod`镜像以及开启`cgo`)
  
然后终端进入到程序目录，运行`go mod download`下载程序依赖

### 2.配置程序

#### 需配置的文件有
- check.go
- flags.go
- init.go

##### check.go
- 21行output变量的网址

##### flags.go
- panel填写面板数据
```go
Panel{
	// 云端文件后缀
	// 对应着output的%s-%s.vf的第二个
	n: "",
	// 验证面板名字
	// 面板页面标题将会显示name中的内容
	name: "",
	// 面板端口
	// 面板的端口
	port: "",
}
```