package http

import (
	"github.com/Kate-liu/GoWebFramework/app/http/module/demo"
	"github.com/Kate-liu/GoWebFramework/framework/gin"
)

// Routes 绑定业务层路由
func Routes(r *gin.Engine) {
	r.Static("/dist/", "./dist/")
	demo.Register(r)
}
