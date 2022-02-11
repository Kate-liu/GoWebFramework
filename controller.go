package main

import (
	"context"
	"fmt"
	"github.com/Kate-liu/GoWebFramework/framework"
	"log"
	"time"
)

func FooControllerHandler(c *framework.Context) error {
	durationCtx, cancel := context.WithTimeout(c.BaseContext(), time.Duration(1*time.Second))
	// 这里记得当所有事情处理结束后调用 cancel，告知 durationCtx 的后续 Context 结束
	defer cancel()

	// 这个 channal 负责通知结束
	finish := make(chan struct{}, 1)
	// 这个 channel 负责通知 panic 异常
	panicChan := make(chan interface{}, 1)
	go func() {
		// 这里增加异常处理
		defer func() {
			if p := recover(); p != nil {
				panicChan <- p
			}
		}()
		// 这里做具体的业务
		// Do real action
		time.Sleep(10 * time.Second)
		c.Json(200, "ok")

		// 新的 goroutine 结束的时候通过一个 finish 通道告知父 goroutine
		finish <- struct{}{}
	}()

	// 请求监听的时候增加锁机制
	select {
	// 监听 panic
	case p := <-panicChan:
		c.WriterMux().Lock()
		defer c.WriterMux().Unlock()
		log.Println(p)
		c.Json(500, "panic")
	// 监听结束事件
	case <-finish:
		fmt.Println("finish")
	// 监听超时事件
	case <-durationCtx.Done():
		c.WriterMux().Lock()
		defer c.WriterMux().Unlock()
		c.Json(500, "time out")
		// 这里记得设置标记为
		c.SetHasTimeout()
	}
	return nil
}
