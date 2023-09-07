package main

import "fmt"

func main() {
	var num int
	num = 100
	fmt.Printf("num:%d,  内存地址：%p", num, &num)
	//和c语言一样，打印地址的时候，需要&取地址符号
	fmt.Println()
	fmt.Printf("\n")
	num = 200
	fmt.Printf("num:%d,  内存地址：%p", num, &num)
	//在这次的执行中，num 是指向同一个的地址
	//但是每一次启动这个程序的时候这个num的地址是变化的
	//因为每次程序的执行结束之后都会销毁这个变量，释放内存空间，所以每次启动这个程序的时候地址都会发生新的变化

}
