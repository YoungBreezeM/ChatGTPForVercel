package api

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func downloadFile(url, destination string) error {
	// 发送GET请求
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// 检查是否成功获取文件信息
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP request failed with status: %s", resp.Status)
	}

	// 创建目标文件
	file, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer file.Close()

	// 将响应体复制到文件
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}

	fmt.Printf("File downloaded to: %s\n", destination)
	return nil
}

func Down(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()

	// 获取特定参数的值
	url := params.Get("url")

	go func() {
		downloadFile(url, "./")
	}()

	w.Write([]byte("OK"))
}
