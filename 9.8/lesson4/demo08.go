package main

import "fmt"

// 递归函数  要有出口
// 尽量不要使用递归函数，因为十分的消耗内存
func main() {

	sum := getSum1(6)
	fmt.Println(sum)

}
func getSum1(n int) int {
	if n == 1 {
		return 1
	}
	return getSum1(n-1) + n

}
