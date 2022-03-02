package http

import (
	"github.com/Kate-liu/GoWebFramework/framework/gin"
	"github.com/Kate-liu/GoWebFramework/provider/demo"
)

func Routes(r *gin.Engine) {

	r.Static("/dist/", "./dist/")

	demo.Register(r)
}
