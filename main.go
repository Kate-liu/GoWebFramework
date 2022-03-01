package main

import (
	"context"
	"github.com/Kate-liu/GoWebFramework/framework/gin"
	"github.com/Kate-liu/GoWebFramework/framework/middleware"
	"github.com/Kate-liu/GoWebFramework/provider/demo"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// core := framework.NewCore()
	//
	// // core中使用use注册中间件
	// // core.Use(
	// // 	middleware.Test1(),
	// // 	middleware.Test2())
	// // group中使用use注册中间件
	// // subjectApi := core.Group("/subject")
	// // subjectApi.Use(middleware.Test3())
	// // core中使用use注册中间件 Recovery
	// core.Use(middleware.Recovery())
	// // core中使用use注册中间件 Cost
	// core.Use(middleware.Cost())
	// // core中使用use注册中间件 Timeout
	// // core.Use(middleware.Timeout(1 * time.Second))

	core := gin.New()
	// 绑定具体的服务
	core.Bind(&demo.DemoServiceProvider{})

	core.Use(gin.Recovery())
	core.Use(middleware.Cost())

	// 设置路由
	registerRouter(core)

	server := &http.Server{
		// 自定义的请求核心处理函数
		Handler: core,
		// 请求监听地址
		Addr: ":8080",
	}

	// 这个 Goroutine 是启动服务的 Goroutine
	go func() {
		server.ListenAndServe()
	}()

	// 当前的 Goroutine 等待信号量
	quit := make(chan os.Signal)
	// 监控信号：SIGINT, SIGTERM, SIGQUIT
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	// 这里会阻塞当前 Goroutine 等待信号
	<-quit

	// 调用Server.Shutdown graceful结束
	if err := server.Shutdown(context.Background()); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
}
