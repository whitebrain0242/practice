package main //要编译称一个可运行程序

import (
	"fmt"
	"io"
	"net"
)

func handleCon(con net.Conn) {
	fmt.Printf("客户端地址是:%s\n", con.RemoteAddr())
	defer con.Close() //管理o生命周期使用defer
	buf := make([]byte, 1024)
	for {
		n, err := con.Read(buf)

		if n > 0 {
			//o成功读取到字节】
			_, Werr := con.Write(buf[:n]) //截取有效个字节
			if Werr != nil {
				fmt.Printf("写回去有问题:%v", Werr)
				return
			}
		}
		if err != nil {
			if err == io.EOF {
				fmt.Printf("客户端：%s,正常退出", con.RemoteAddr())
			} else {
				fmt.Printf("客户端,%s 有问题：%v", con.RemoteAddr(), err)
			}
			break
		}
	}

	fmt.Println("连接已关闭")
}
func main() { //最终执行入口
	//首先开始监听
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Printf("监听失败:%v\n", err) //使用Println的花是4识别不了参数的i直接输出了
		return
	}
	//已经开始监听，接下来直接开始管理生命周期，accept
	defer listener.Close() //压入栈中，然后main结束之前执行
	//开始接受连接
	for {
		con, err := listener.Accept()
		if err != nil {
			fmt.Printf("连接失败:%v\n", err)
			return
		}
		handleCon(con)
	}

}
