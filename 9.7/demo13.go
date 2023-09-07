package main

import "fmt"

func main() {
	//数据类型的转换

	//go语言中所有的转换必须都是显式的！！！
	//go语言没有隐式类型转换
	a := 3
	b := 5.0

	c := float64(a)
	d := int(b)
	//整形是不能够转化成布尔型的
	//e:=bool(a)   是错的
	//int ->float64!!!

	//数据转化要注意范围

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
	//默认的转换也是转换成float64
	fmt.Printf("%T\n", d)

}
