package main

import (
	"io"
	"net"
	"testing"
)

func TestHandleConnEcho(t *testing.T) {
	//i先在内存中设置一个通道
	serverConn, clientConn := net.Pipe()
	defer clientConn.Close()

	//副线程运作服务器，主线称运作客户端
	go handleCon(serverConn)

	//写入数据
	want := "hello"
	n, Werr := io.WriteString(clientConn, want)
	if Werr != nil {
		t.Fatalf("i写入失败：%v", Werr)
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
}
