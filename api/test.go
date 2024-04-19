package api

import (
	"cfv/service"
	"net/http"
)

func Test(w http.ResponseWriter, r *http.Request) {
	go service.Start()

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("请求已收到，正在处理中"))

}
