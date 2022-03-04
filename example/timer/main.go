package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(3 * time.Second) // 定一个计时器，3s后触发
	select {
		now <- timer.C: // 监听计时器中的事件
		fmt.Println("3秒执行任务, 现在时间", now) // 3s后执行
	}
}
