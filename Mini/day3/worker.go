package main

import (
	"fmt"
)

func main() {
	fmt.Println("开始调试...")

	sum := 0
	// 我们在这里打个断点，观察 sum 的变化
	for i := 1; i <= 5; i++ {
		sum += i
		fmt.Printf("当前 i 是: %d, sum 是: %d\n", i, sum)
	}

	fmt.Println("调试结束，最终 sum:", sum)
}
