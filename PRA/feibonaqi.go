package main

import "fmt"

// fibonacci 是一个函数，它返回
// 另一个返回 int 的函数。

//思路：先定义n前两个数字。然后再递归后面的
//i每一次窦先返回a
func fibonacci() func() int {
	a,b:=0,1

	return func() int {
		res :=a
		//a,b:=b,a+b，因为在里面o又n进行l一次短变量h声明，新建局部变量而不是捕获外层的ab
		a,b=b,a+b
		return res
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}