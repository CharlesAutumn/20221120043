package main

import "fmt"

func main() {
	//闭包    由于函数回调   有函数回调的 也有这种特性---闭包  JavaScript

	r1 := increment()
	fmt.Println(r1)

	v1 := r1()
	fmt.Println(v1)

	//0xbe3b80
	//1

	v2 := r1()
	fmt.Println(v2)
	fmt.Println(r1())
	fmt.Println(r1())
	fmt.Println(r1())

	r2 := increment()
	fmt.Println(r2())
	fmt.Println(r1())
	//0xe83d80
	//1
	//2
	//3
	//4
	//5
	//1   因为重新调用了increment 重置了外层的局部变量
	//6   为什么还是6  因为内层局部变量进行了保存   在开发中，尽量少使用中这种闭包结构  但是由于这种闭包特性，也可以解决一些特定的问题

}

// 自增
func increment() func() int {
	//局部变量
	i := 0
	//定义一个匿名函数，给变量自增并返回
	fun := func() int { //内层函数，没有执行
		i++
		return i
	}
	return fun
}
