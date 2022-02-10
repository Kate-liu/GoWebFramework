package main

import "net/http"

func main() {
	// 对文件读取的 HTTP 服务
	fs := http.FileServer(http.Dir("/home/bob/static"))
	http.Handle("/static/", http.StripPrefix("/static", fs))
}
