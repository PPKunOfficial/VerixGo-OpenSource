package Core_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/PPKunOfficial/VerixGo/Core"
)

func TestDownloadFile(t *testing.T) {
	// 创建一个 HTTP 服务器并设置处理函数
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/octet-stream")
		_, err := w.Write([]byte("Test file content"))
		if err != nil {
			return
		}
	}))
	defer server.Close()

	// 下载文件
	url := server.URL
	fileContent, err := Core.DownloadFile(url)
	if err != nil {
		t.Errorf("DownloadFile() failed: %v", err)
	}

	// 验证下载的文件内容是否正确
	expectedContent := "Test file content"
	if string(fileContent) != expectedContent {
		t.Errorf("Downloaded file content does not match: expected %s, got %s", expectedContent, string(fileContent))
	}
}
