package main

import "fmt"

func main() {

	//go中的字符串是一个一个字节拼起来的

	str := "hhhhh,h"
	fmt.Println(str)
	//获取字符串的长度 len
	fmt.Println("字符串的长度为：", len(str)) //7
	//注意：这里是用逗号，将”“内的内容与数值进行连接的

	//获取指定的字节
	fmt.Println("打印字节：", str[0]) //104  h 的ASCII编码

	for i := 0; i < len(str); i++ {
		//fmt.Println(str[i])  //输出的是数字对的编码
		fmt.Printf("%c\n", str[i])

	}

	//有一种特殊的for循环，由于遍历数组和切片
	for i, v := range str {
		fmt.Print(i)
		fmt.Printf("%c", v)

		//0h1h2h3h4h5,6h

	}

	//str[2] = 'A'
	//cannot assign to str[2] (neither addressable nor a map index expression)
	//string  是不能被修改的

}
