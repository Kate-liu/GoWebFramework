package main

import "github.com/Kate-liu/GoWebFramework/framework"

func registerRouter(core *framework.Core) {
	// 设置控制器
	core.Get("foo", FooControllerHandler)
}
