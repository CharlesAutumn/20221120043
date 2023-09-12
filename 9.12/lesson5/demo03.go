package main

import "fmt"

func main() {
	var arr01 []int //声明切片  打印一个空数组  nil

	fmt.Println(arr01 == nil) //true  切片的默认值就是nil

	//arr01[0] = 12没有办法通过这种方式对切片扩容
	fmt.Printf("%v - %T - %v", arr01, arr01, len(arr01))

	var arr02 = []int{1, 2, 3, 4, 5}
	fmt.Println(arr02)

	var arr03 = []int{1: 3, 5: 6}
	fmt.Println(arr03)
}
