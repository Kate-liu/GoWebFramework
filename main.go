package main

import (
	"github.com/Kate-liu/GoWebFramework/framework"
	"net/http"
)

func main() {
	core := framework.NewCore()

	// 设置路由
	registerRouter(core)

	server := &http.Server{
		// 自定义的请求核心处理函数
		Handler: core,
		// 请求监听地址
		Addr: ":8080",
	}

	server.ListenAndServe()
}
