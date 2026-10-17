package main //要编译称一个可运行程序

import (
	"fmt"
	"net"
)

func  main(){//最终执行入口
	listener,err:=net.Listen("tcp","127.0.0.1:8000")	
	if err !=nil{
		fmt.Printf("监听失败:%v\n",err)//使用Println的花是4识别不了参数的i直接输出了
		return
	}
	//已经开始监听，接下来直接开始管理生命周期，accept
	defer listener.Close()//压入栈中，然后main结束之前执行
	con,err:=listener.Accept()
	if err !=nil{
		fmt.Printf("连接失败:%v\n",err)
		return
	}
	fmt.Printf("e客户端地址是:%s\n",con.RemoteAddr())
	con.Close()
	fmt.Println("连接已关闭")
}

