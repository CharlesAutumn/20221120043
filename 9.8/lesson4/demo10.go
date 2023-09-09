package main

import "fmt"

// func() 本身就是一个数据类型
func main() {

	fmt.Printf("%T", f10)
	fmt.Println()
	//func() 本身就是数据类型
	// 打印整个函数的本质的一些东西，是不要加括号的

	//f10如果不加括号的话就是一个变量

	var f5 func(int, int) int
	f5 = f10
	c := f5(1, 2)
	fmt.Println(c)

	fmt.Println(f5)
	fmt.Println(f10)
	//0x735600
	//0x735600
	//指向了同一个地址

	//函数在go中是一种复合类型，可以看作是一种特殊的变量
	//函数名（）:调用返回结果
	//函数：指向函数体的内存地址，一种特殊类型的指针变量

}
func f10(a, b int) int { //此时func的类型是 func(int, int) int
	return a + b

	//可以发现函数的类型就是函数定义的本身
}
