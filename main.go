package main

import (
	"github.com/Kate-liu/GoWebFramework/app/console"
	"github.com/Kate-liu/GoWebFramework/app/http"
	"github.com/Kate-liu/GoWebFramework/framework"
	"github.com/Kate-liu/GoWebFramework/framework/provider/app"
	"github.com/Kate-liu/GoWebFramework/framework/provider/kernel"
)

func main() {
	// 初始化服务容器
	container := framework.NewHadeContainer()

	// 绑定 App 服务提供者
	container.Bind(&app.HadeAppProvider{})

	// 后续初始化需要绑定的服务提供者...
	// 将 HTTP 引擎初始化,并且作为服务提供者绑定到服务容器中
	if engine, err := http.NewHttpEngine(); err == nil {
		container.Bind(&kernel.HadeKernelProvider{HttpEngine: engine})
	}

	// 运行 root 命令
	console.RunCommand(container)

}
