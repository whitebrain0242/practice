package main

import (
	"fmt"
)

// // 思路：  现在要定义一个动物e叫声的接口
// type Speaker interface {
// 	Speak()
// }//接口也是一个数据类型，

// //只要数据类型实现了接口中的所有方法，不用申请就可以直接放进去——————所以说o接口是i隐士实现的
// type Dog struct{}
// func (d Dog)Speak(){
// 	fmt.Println("汪汪汪")
// }
// type Cat struct {}
// func (c Cat)Speak (){
// 	fmt.Println("喵喵瞄")
// }

// // 然后定义一个关于最外部 的调用n函数
// func MakeNoise(s Speaker){
// 	s.Speak()
// }

// func main(){
// 	//先定义一个盒子
// 	//var box Speaker——接口n变量保存i一对信息：数据i类型和n底层的值，其中及时值是nil依然可以进行正常的调用，因为数据类型不是空的——只有当 数据类型和值都是空的i 时候，接口才是空的
// 	//——————如果当a两个值都是nil的时候a那么，此时再调用方法的花，可能就会报错panic,（严重错误，程序直接停止）
// 	//把狗和猫类型加进去
// 	// box=Dog{}
// 	// MakeNoise(box)
// 	// box=Cat{}
// 	// MakeNoise(box)
// 	//其实可以直接传输
// 	MakeNoise(Dog{})
// 	MakeNoise(Cat{})
// }

// //空接口
// func printAny(v any){
// 	fmt.Printf("值：%v,类型 ：%T\n",v,v)
// }
// //空接口可以表示任意类型，接受任何类型
// func main(){
// 	printAny(100)//值：100,类型：int
// 	printAny("golang")//值：golang,类型：string
// 	printAny([]int{1,2,3})//值：[1,2,3],类型：【】int
// }

// //类型断言———从变量中u取出具体类型的值，来看和目标类型是否一致
// func main(){
// 	var i any="golang"
//     //第一各s是值，第二各是bool类型是否一致
// 	//val, ok := 接口变量.(目标类型)
// 	s,ok:=i.(string)
// 	if ok{
// 		fmt.Println("是字符串，值是",s)
// 	}else {
// 		fmt.Println("不是字符串")
// 	}
// 	num,ok:=i.(int)
// 	if ok {
// 		fmt.Println("是整数，值为" ,num)
// 	}else{
// 		fmt.Println("不是整数，值是",num)
// 	}
// }

// //type switch————专门一次性进行 多种类型 判断
// func check(i any){
// 	switch v :=i.(type){
// 	case int:
// 		fmt.Println("zhengshu",v)

// 	case float64:
// 		fmt.Println("fudianshu",v)

// 	default:
// 		fmt.Println(wu)
// 	}

// }

//Stringeru接口
type Ver struct{X,Y int}

//加上srting方法,那么现在已经把这个类型加入了 接口中
func (v Ver)String()string{
	 return fmt.Sprintf("X=%d,Y=%d",v.X,v.Y)
}
func main(){
	//最后的结果不是1,2.而是X=1,Y=2
	//其实就是说只要你的类型实现了 `String() string` 方法，`fmt` 系列打印函数（Println/Printf）输出这个变量时，就自动调用你写的 `String()`，不再打印默认原始格式。
	fmt.Println(Ver{1,2})
}