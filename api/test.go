package api

import (
	"fmt"
	"net/http"
	"time"
)

func Test(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	// 设置跨域访问
	w.Header().Set("Access-Control-Allow-Origin", "*")

	for i := 0; i < 10; i++ {

		fmt.Fprintf(w, "data: Initial data\n\n")

		time.Sleep(time.Second * 1)
	}
}
