package main

import "fmt"

func main() {
	//函数的调用
	printInfo()

	myPrint("hhhhhhh")

	x, y := swap("第一个", "第二个")
	fmt.Println(x, y)
}

func printInfo() {

	fmt.Println("printInfo")

}

func myPrint(msg string) {

	fmt.Println(msg)

}

func add2(a, b int) int { //!!!!注意，这里int   没有加（）可以的

	c := a + b
	return c
}

func swap(x, y string) (string, string) { //这里就是返回值有多个的时候价格括号，理所应当
	return y, x
}
