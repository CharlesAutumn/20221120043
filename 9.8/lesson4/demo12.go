package main

import (
	"fmt"
)

func main() {
	r1 := add(1, 2)
	fmt.Println(r1)

	r2 := oper(3, 4, add)
	fmt.Println(r2)

	r3 := oper(3, 4, sub)
	fmt.Println(r3)

	r4 := oper(8, 4, func(a int, b int) int { //这是匿名函数！！！！
		if b == 0 {
			fmt.Println("除数不能为零")
		}
		return a / b

	})
	fmt.Println(r4)

}

func oper(a, b int, fun func(int, int) int) int { //这个就叫高阶函数

	r := fun(a, b)
	return r

}

func add(a, b int) int { //这个add就叫做回调函数
	return a + b

}

func sub(a, b int) int {
	return a - b
}
