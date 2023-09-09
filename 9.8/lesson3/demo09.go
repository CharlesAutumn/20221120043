package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			break //!!!!整个！！！！直接结束掉整个循环
		}
		fmt.Println(i)

	}

	for i := 0; i < 10; i++ {
		if i == 5 {
			continue //!!!!!!当次!!!!!跳过当次循环，继续执行
		}
		fmt.Println(i)

	}
	fmt.Println()

}
