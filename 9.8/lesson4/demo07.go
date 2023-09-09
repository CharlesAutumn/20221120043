package main

import "fmt"

// 全局变量
var num int = 100

func main() {

	temp := 100
	//if语句  for 语句 定义的一次性变量为局部变量
	if b := 1; b < 10 {
		//语句内的局部变量
		temp := 50
		fmt.Println(temp) //局部变量，就近原则
		fmt.Println(b)
	}
	fmt.Println(temp)
	//fmt.Println(b)  直接报红

	//50
	//1
	//100

	fmt.Println(num)
	f1()

}

func f1() {
	a := 1
	fmt.Println(a)
	num := 33 //局部优先
	fmt.Println(num)
}
func f2() {
	//fmt.Println(a)  就是理解一下局部变量
}
