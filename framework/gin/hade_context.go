// Package gin Copyright 2021 jianfengye.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package gin

import (
	"context"
)

func (ctx *Context) BaseContext() context.Context {
	return ctx.Request.Context()
}

//  context 实现 container 的几个封装

// Make 实现 make 的封装
func (ctx *Context) Make(key string) (interface{}, error) {
	return ctx.container.Make(key)
}

// MustMake 实现 mustMake 的封装
func (ctx *Context) MustMake(key string) interface{} {
	return ctx.container.MustMake(key)
}

// MakeNew 实现 makenew 的封装
func (ctx *Context) MakeNew(key string, params []interface{}) (interface{}, error) {
	return ctx.container.MakeNew(key, params)
}
