package main

import "fmt"

//T是类型参数，T只能是int或者是float64
func  add [T int|float64](a,b T)T{
	return a+b
}
func main(){
	//自动判断类型
	fmt.Println(add (1.5,2))
	boxInt :=Box[int]{val: 100}
}
//T 就是一各泛型，其中的val可以保存任何类型
//这样的花个定义，就可以装不同的类型了
type Box[T any] struct{
	val T
}