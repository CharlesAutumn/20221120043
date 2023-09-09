package main

import "fmt"

func main() {
	getSum("哈哈哈哈", 1, 2, 3, 4, 5)
}

// 可变参数，就是传入的参数可以是多个    ...int
// ！！！可变参数要定义再最后面
// 一个参数列表中最多只能有一个可变参数
func getSum(msg string, nums ...int) {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]

	}
	fmt.Println(sum)

}
