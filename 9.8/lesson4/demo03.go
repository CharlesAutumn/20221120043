package main

import "fmt"

func main() {

	k := max(1, 2)
	fmt.Println(k)
	//习惯做好
}

func max(num1, num2 int) int {
	var result int
	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}
