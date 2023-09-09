package main

import "fmt"

func main() {
	//打印99乘法表
	for j := 1; j <= 9; j++ {
		for i := 1; i <= j; i++ {
			fmt.Printf("%d*%d=%d \t", i, j, i*9)

		}
		fmt.Println()

	}
}
