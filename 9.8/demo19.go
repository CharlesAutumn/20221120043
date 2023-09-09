package main

import "fmt"

func main() {
	var x int
	var y float64
	fmt.Println("请输入两个数")

	fmt.Scanln(&x, &y)

	fmt.Println(x, y)

}
