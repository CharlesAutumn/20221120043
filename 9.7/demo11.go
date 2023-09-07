package main

import "fmt"

func main() {
	var age int = 18
	//uint 表示0以上的，包扣0
	fmt.Println(age)
	var money float64 = 3.14
	//默认打印6位小数(小数点后六位）
	fmt.Printf("%.1", money) //.1  这个保留小数的操作，是四舍五入的

	//byte  == uint8
	//rune == int32
	//int == int64

	//浮点数主要使用float64，float默认64
}
