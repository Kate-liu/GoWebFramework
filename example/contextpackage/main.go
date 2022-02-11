package contextpackage

import (
	"github.com/Kate-liu/GoWebFramework/framework"
	"net/http"
)

// 控制器
func Foo2(ctx *framework.Context) error {
	obj := map[string]interface{}{
		"data": nil,
	}

	// 从请求体中获取参数
	fooInt := ctx.FormInt("foo", 10)

	// 构建返回结构
	obj["data"] = fooInt

	// 输出返回结构
	return ctx.Json(http.StatusOK, obj)
}

// 封装调用
func Foo3(ctx *framework.Context) error {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return rdb.Set(ctx, "key", "value", 0).Err()
}
