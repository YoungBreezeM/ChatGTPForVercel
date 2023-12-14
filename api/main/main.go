package main

import (
	"cfv/api"
	"fmt"
)

func main() {
	if err := api.DownloadFile("http://124.221.78.219:28158/bore", "./bore"); err != nil {
		fmt.Println(err)
	}
}
