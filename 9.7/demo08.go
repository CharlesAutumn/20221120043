package main

import "fmt"

func main() {
	const URL string = "www.baidu.com" //显式定义
	const URL2 = "www.baidu.com"       //隐式定义
	var a = "shibushi"                 //var a string = "shibushi"

	const d, f, g = 13, "hhhh", 2

	fmt.Println(d, f, g)
	b := "666666"
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(URL)
	fmt.Println(URL2)
	//URL = "www.baidu.com"  这个run不起来，常量是不可以被改变

}
