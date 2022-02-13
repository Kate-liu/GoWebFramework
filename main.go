package main

import (
	"github.com/Kate-liu/GoWebFramework/framework"
	"github.com/Kate-liu/GoWebFramework/framework/middleware"
	"net/http"
)

func main() {
	core := framework.NewCore()

	// core中使用use注册中间件
	// core.Use(
	// 	middleware.Test1(),
	// 	middleware.Test2())
	// group中使用use注册中间件
	// subjectApi := core.Group("/subject")
	// subjectApi.Use(middleware.Test3())
	// core中使用use注册中间件 Recovery
	core.Use(middleware.Recovery())
	// core中使用use注册中间件 Cost
	core.Use(middleware.Cost())
	// core中使用use注册中间件 Timeout
	// core.Use(middleware.Timeout(1 * time.Second))

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
