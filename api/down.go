package api

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func DownloadFile(url, destination string) error {
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
	var ch chan struct{}

	go func() {
		fmt.Println("start download")
		if err := DownloadFile(url, "./bore"); err != nil {
			fmt.Println(err)
		}
		fmt.Println("end download")
		ch <- struct{}{}

	}()

	w.Write([]byte("OK"))

	<-ch
}
