package main

import (
	"io"
	"net"
	"testing"
	"time"
)

func TestHandleConnEcho(t *testing.T) {
	t.Run("normal hello message", func(t *testing.T) {
		//先在内存中设置一个通道
		serverConn, clientConn := net.Pipe()
		defer clientConn.Close()

		//副线程运作服务器，主线称运作客户端
		go handleCon(serverConn)

		//写入数据
		want := "hello"
		n, Werr := io.WriteString(clientConn, want)
		if Werr != nil {
			t.Fatalf("写入失败：%v", Werr)
		}
		if n != len(want) {
			t.Fatalf("字数不对，写入：%d, 想要：%d", n, len(want))
		}
		//回显数据
		got := make([]byte, len(want))
		_, err := io.ReadFull(clientConn, got)
		if err != nil {
			t.Fatalf("读取失败：%v", err)
		}
		if string(got) != want {
			t.Fatalf("回显内容错误")
		}
	})

	t.Run("close", func(t *testing.T) {
		//先在内存中设置一个通道
		serverConn, clientConn := net.Pipe()
		defer clientConn.Close()

		done := make(chan struct{}) //接受断开时信息
		//新开e一个县城运行服务端
		go func() {
			handleCon(serverConn)
			close(done)
		}() //这个括号是调用的意思
		//关闭 客户端，观察服务端是否有信息
		clientConn.Close()
		//同时等待多个事件，那个先发生就先i执行谁
		select {
		case <-done: //接受EOF信息
		case <-time.After(time.Second):
			t.Fatalf("handlecon 没有断开时返回")
		}
	})

}
