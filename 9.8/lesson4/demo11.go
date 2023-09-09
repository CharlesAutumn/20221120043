package main

import (
	"fmt"
)

func main() {
	fmt.Println()
	f1()
	f2 := f1 //函数本身也是一个变量
	f2()

	//匿名函数
	f3 := func() {
		fmt.Println("我是f3函数")
	}
	f3()

	//匿名函数自己 调用自己
	func() {
		fmt.Println("我是f4函数")
	}()

	r1 := func(a, b int) int {
		fmt.Println(a, b)
		return a + b

	}(1, 2)
	fmt.Println(r1)

}

func f1() {
	fmt.Println("我是f1函数")

}
