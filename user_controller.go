package main

import (
	"github.com/Kate-liu/GoWebFramework/framework"
	"time"
)

func UserLoginController(c *framework.Context) error {
	foo, _ := c.QueryString("foo", "def")
	// 等待10s才结束执行
	time.Sleep(10 * time.Second)
	// 输出结果，打印控制器名字
	c.SetOkStatus().Json("ok, UserLoginController: " + foo)
	return nil
}
