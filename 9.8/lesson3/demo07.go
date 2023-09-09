package main

import "fmt"

func main() {

	for j := 0; j < 5; j++ {
		//有快捷操作 fori 再改变量
		for i := 1; i <= 5; i++ {
			fmt.Print("* ")
		}
		fmt.Println()

	}

}
