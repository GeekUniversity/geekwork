package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
)

// getEnv 读取当前系统的环境变量中的 VERSION 配置，并写入 response header
func getEnv() string {
	VERSION := os.Getenv("VERSION")
	return VERSION
}

// healthz 当访问 localhost/healthz 时，应返回 200
func healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	// 记录访问信息
	stdout(r)
}

// stdout Server 端记录访问日志: 包括客户端IP、HTTP返回码, 输出到 server 端的标准输出
func stdout(r *http.Request) {
	// 创建一个 recorder 用来保存响应信息
	recorder := httptest.NewRecorder()
	recorder.Flush()
	// 记录日志
	addr := r.RemoteAddr
	httpStatusCode := recorder.Code
	logData := fmt.Sprintf("客户端IP=%v HTTP返回码=%v\n", addr, httpStatusCode)
	// 标准输出
	fmt.Fprintf(os.Stdout, logData)
}

func errLog(r *http.Request, code int) {
	// 记录日志
	addr := r.RemoteAddr
	logData := fmt.Sprintf("客户端IP=%v HTTP返回码=%v\n", addr, code)
	// 标准输出
	fmt.Fprintf(os.Stdout, logData)
}

func writeRespHeader(w http.ResponseWriter, r *http.Request) {
	headers := r.Header
	// 遍历Headers, 将每一个req header内容写入到response header中
	for key, val := range headers {
		w.Header().Set(key, val[0])
	}
	// 响应头添加VERSION
	VERSION := getEnv()
	w.Header().Set("VERSION", VERSION)
	// 记录访问信息
	stdout(r)
}

// Handler 接收客户端 request，并将 request 中带的 header 写入 response header
func Handler(w http.ResponseWriter, r *http.Request) {
	// 在http包使用的时候，注册了/这个根路径的模式处理，浏览器会自动的请求favicon.ico, 一定要注意处理 ，否则会出现两次请求
	if r.URL.RequestURI() == "/favicon.ico" {
		return
	}
	// 获取请求路由, 根据不同路由匹配不同的方法
	path := r.URL.Path
	if path == "/" {
		writeRespHeader(w, r)
	} else if path == "/healthz" {
		healthz(w, r)
	} else {
		w.WriteHeader(500)
		// 记录访问信息
		errLog(r, 500)
	}
}

func main() {
	// 设置处理函数
	http.HandleFunc("/", Handler)
	//http.HandleFunc("/healthz", healthz)
	// 监听80端口
	err := http.ListenAndServe("0.0.0.0:80", nil)
	if err != nil {
		fmt.Printf("服务端启动失败, err=%v\n", err)
		return
	}
}
