package main

import "github.com/Kate-liu/GoWebFramework/framework"

func UserLoginController(c *framework.Context) error {
	// 打印控制器名字
	c.SetOkStatus().Json("ok, UserLoginController")
	return nil
}
