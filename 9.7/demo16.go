package main

import "fmt"

func main() {
	//逻辑运算符。。。。
	//   &&    都为真才真
	//   ||    都为假才为假
	//   ！
	var a bool = true
	var b bool = true
	if a == true && b == true {
		fmt.Println(a && b)
	}

	fmt.Println()

}
