package api

import (
	"cfv/service"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ChatGTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	// 设置跨域访问
	w.Header().Set("Access-Control-Allow-Origin", "*")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	chat := service.Chat{}
	req := service.ChatRequest{}
	if err := json.Unmarshal(body, &req); err != nil {
		w.Write([]byte("fail"))
		return
	}

	chat.Chat(&req, w)

}
