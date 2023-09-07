package main

import "fmt"

func main() {
	//a=100
	//b=200
	//temp=0
	//1、temp=a
	//2、a=b
	//3、b=temp(a)

	var a int = 100
	var b int = 200

	b, a = a, b //特点！
	fmt.Println(a, b)

}
