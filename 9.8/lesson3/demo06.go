package main

import "fmt"

func main() {
	//注意，在go语言中，循环只有for，没有while循环
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	//简化，爽
	j := 1
	for j <= 10 {
		fmt.Println(j)
		j++
	}

	//再简化
	//这是无限循环
	k := 1
	for {
		fmt.Println(k)
		k++
	}

}
