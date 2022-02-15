package main

import (
	"github.com/Kate-liu/GoWebFramework/framework/gin"
	"time"
)

func UserLoginController(c *gin.Context) {
	foo, _ := c.DefaultQueryString("foo", "def")
	// 等待10s才结束执行
	time.Sleep(10 * time.Second)
	// 输出结果，打印控制器名字
	c.ISetOkStatus().IJson("ok, UserLoginController: " + foo)
}
