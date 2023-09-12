package main

import "fmt"

func main() {
	var str = "你好golang" //首先，普通for循环无法打印汉字，要使用range

	for k, v := range str {
		//fmt.Println(k, v)
		//0 20320  注意是数字
		//3 22909
		//6 103
		//7 111
		//8 108
		//9 97
		//10 110
		//11 103
		fmt.Printf("%v %c\n", k, v)
		//0 你
		//3 好
		//6 g
		//7 o
		//8 l
		//9 a
		//10 n
		//11 g

	}

	var arr = []string{"php", "java", "node", "golang"}
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
		//php
		//java
		//node
		//golang
	}

	for _, val := range arr {
		fmt.Println(val)
		//php
		//java
		//node
		//golang

	}
}
