package main

import (
	"fmt"
	"net/http"
)

func helloHander(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "hello, web")
}
func main() {
	//注册路由与处理函数
	http.HandleFunc(".", helloHander)
	//启动服务器并监听
	http.ListenAndServe()
}
