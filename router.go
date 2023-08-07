package main

import (
	"html/template"
	"io/fs"
	"net/http"

	"github.com/PPKunOfficial/VerixGo/public"
	"github.com/gin-gonic/gin"
)

func HttpServer() {
	r := gin.New()
	err := r.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		return
	}

	templatePaths, err := fs.Glob(public.Public, "dist/*.html")
	if err != nil {
		log.Errorf("解析模板文件失败:%s", err)
	}

	tmpl := template.Must(template.ParseFS(public.Public, templatePaths...))

	r.SetHTMLTemplate(tmpl)

	// 定义路由
	// 首页
	r.GET("/", func(c *gin.Context) {
		if vLocal.Verify {
			c.HTML(http.StatusOK, "already_login.html", gin.H{
				"title": panel.name,
			})
		} else {
			c.HTML(http.StatusOK, "index.html", gin.H{
				"title": panel.name,
			})
		}
	})

	// 登录验证API
	r.POST("/api/user/login", func(c *gin.Context) {
		st := CheckOld()
		c.JSON(http.StatusOK, gin.H{
			"status": st,
		})
	})

	err = r.Run(":" + panel.port)
	if err != nil {
		return
	}
}
func DebugHttp() {
	if !Debug {
		return
	}
	r := gin.New()
	err := r.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		return
	}

	// 定义路由
	r.GET("/", func(c *gin.Context) {
		// 读取文件内容
		content := logByteBuffer.String()
		if err != nil {
			c.String(http.StatusInternalServerError, "Failed to read file")
			return
		}

		// 返回文件内容
		c.String(http.StatusOK, string(content))
	})

	err = r.Run(":8080")
	if err != nil {
		return
	}
}
