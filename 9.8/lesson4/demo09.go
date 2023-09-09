package main

import "fmt"

func main() {

	//defer

	f5("1")
	f5("2")
	defer f5("3") //先进后出，栈形式   从下往上输出defer
	f5("4")
	defer f5("5")

	fmt.Println("---------------")

	a := 10
	fmt.Println("a=", a)
	defer f6(a) //10  !!!!!!参数已经传递进去了，最后执行
	a++
	fmt.Println("end a = ", a)

}

func f5(s string) {
	fmt.Println(s)

}
func f6(s int) {
	fmt.Println(s)

}
