package main

//maina代表是可执行程序,必须有 main入口函数

import (
	"strings"
)

// func WordCount(s string) map[string]int {
// 	//chai
// 	words :=strings.Fields(s)

// 	//映射,必须使用mkae初始化才能存值
// 	cntMap:=make(map[string]int)

// 	//遍历切片，——代表忽略地几个元素的索引
// 	for _,word:=range words{
// 		_,ok :=cntMap[word]//map有两个返回值：i只有i 一个：值‘’‘两个”i值加是否存在
// 		if ok{
// 			cntMap[word]=cntMap[word]+1
// 		}else{
// 			cntMap[word]=1
// 		}
// 	}
// 	return map[string]int{"x": 1}
// }

// func main() {
// 	wc.Test(WordCount)
// }
// 思路：1,先把字符串s中的每一个单词分开，保存在一个容器中，可以是数组把
// 2. 然后 遍历这个每一个单词，创建一个 映射，先看映射里面有没有这个单词的键
//    如过有，那么值加一，如果没有，那么新加一个键，u并且计数1
// 3.最后返回这个映射

//进阶
func WordCount(s string) map[string]int {
	//chai
	words :=strings.Fields(s)
	
	//映射,必须使用mkae初始化才能存值
	cntMap:=make(map[string]int)

	//遍历切片，——代表忽略地几个元素的索引
	for _,word:=range words{
		cntMap[word]++
	}
	return map[string]int{"x": 1}
}