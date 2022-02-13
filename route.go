package main

import (
	"github.com/Kate-liu/GoWebFramework/framework"
	"github.com/Kate-liu/GoWebFramework/framework/middleware"
)

// 注册路由规则
func registerRouter(core *framework.Core) {
	// 需求1+2:HTTP方法+静态路由匹配
	// core.Get("/user/login", UserLoginController)
	// 在核心业务逻辑 UserLoginController 之外，封装一层 TimeoutHandler
	// core.Get("/user/login", framework.TimeoutHandler(UserLoginController, time.Second))
	// 在core中使用middleware.Test3() 为单个路由增加中间件
	core.Get("/user/login", middleware.Test3(), UserLoginController)

	// 需求3:批量通用前缀
	subjectApi := core.Group("/subject")
	{
		// 需求4:动态路由
		subjectApi.Delete("/:id", SubjectDelController)
		subjectApi.Put("/:id", SubjectUpdateController)
		subjectApi.Get("/list/all", SubjectListController)
		// 在group中使用middleware.Test3() 为单个路由增加中间件
		// subjectApi.Get("/:id", SubjectGetController)
		subjectApi.Get("/:id", middleware.Test3(), SubjectGetController)

		// 思考题：批量通用前缀
		subjectInnerApi := subjectApi.Group("/info")
		{
			subjectInnerApi.Get("/name", SubjectNameController)
		}
	}

}
