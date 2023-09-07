package main

import "fmt"

func test() (int, int) { //这个地方要注意与c、java 的区别

	return 100, 200

}

func main() {
	a, _ := test()
	_, b := test()
	//使用匿名变量
	//     _  	会被废弃掉，其实使用匿名变量的原因是因为，
	//对于某个函数，返回值有两个，在main 中必须要有两变量去接收，
	//但是其中有变量的值没有必要区接收，这个时候就可以使用匿名变量
	//使得接受出来的值直接废弃掉！！！注意是接受出来的值，如果再次调用这个函数，100，200还是存在的
	fmt.Println(a)
	//注意在go中定义变量之后不去使用会报错的
	fmt.Println(b)

	//补充说明：关于全局变量、局部变量的使用
	//在main函数等函数以外定义的变量--全局变量
	//在函数内部定义的变量--局部变量
	//有一种情况----在定义完全局变量之后，
	//在函数内部还可以对全局变量同名的变量进行局部的定义，
	//在该函数内输出的变量值为局部定义的值

	//注意作用域就行

}