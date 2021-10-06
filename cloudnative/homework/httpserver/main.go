package main

import (
	"fmt"
	"net/http"
)

func Handler() {

}

func main() {
	// 监听80端口
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		fmt.Printf("服务端启动失败, err=%v\n", err)
		return
	}
	//
}
