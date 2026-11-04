package main

import (
	"fmt"
	"log"
	"net/http"
)

// w是响应写入器,通过这个把内容发给客户端,r表示这一次的客户端请求
func helloHander(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	fmt.Println(r.URL.Path)
	fmt.Println(r.RemoteAddr)

	fmt.Fprintln(w, "hello from provider") //把内容写入w中发给客户端
}
func main() {
	//注册路由:URL匹配规则(路由模式),
	http.HandleFunc("/hello", helloHander)
	//启动服务器并监听
	http.ListenAndServe("127.0.0.1:8080", nil)
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}
