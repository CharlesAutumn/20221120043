package main

import "fmt"

func main() {

	//算术运算符
	//加减乘除，%，++，--
	var a int = 10
	var b int = 3
	//var c int
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a % b)
	a++
	fmt.Println(a) //直接输出a++是不好使的，需要先a++再输出a
	a--
	fmt.Println(a)

}
