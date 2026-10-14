package main

import "fmt"

// 思路：  现在要定义一个动物e叫声的接口
type Speaker interface {
	Speak()
}//接口也是一个数据类型

//只要数据类型实现了接口中的所有方法，那么就可以归属于接口
type Dog struct{}
func (d Dog)Speak(){
	fmt.Println("汪汪汪")
}
type Cat struct {}
func (c Cat)Speak (){
	fmt.Println("喵喵瞄")
}

// 然后定义一个关于最外部 的调用n函数
func MakeNoise(s Speaker){
	s.Speak()
}

func main(){
	//先定义一个盒子
	//var box Speaker
	//把狗和猫类型加进去
	// box=Dog{}
	// MakeNoise(box)
	// box=Cat{}
	// MakeNoise(box)
	//其实可以直接传输
	MakeNoise(Dog{})
	MakeNoise(Cat{})
}