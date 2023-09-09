package main

import "fmt"

func main() {
	//值传递：传递的是数据的副本，修改数据，对于原始的数据没有影响
	//值类型的数据，默认都是值传递，基础类型，array struct
	//定义数组   【个数】 类型
	arr := [4]int{1, 2, 3, 4}
	upDate(arr)
	//实际上是进行了一个拷贝动作

	fmt.Println(arr)
	//[1 2 3 4]
	//第一个，用中括号[]打印出来的都是数组
	//第二个，ln了，但是没有换行
	//第三个，调用函数后，回到主程序，原数组是不变的
	//实际上，arr 与 arr2 是两个数据了，因为执行了拷贝动作

	//引用传递
}

func upDate(arr2 [4]int) {
	fmt.Println(arr2)
	arr2[0] = 100
	fmt.Println(arr2)

	//[1 2 3 4]
	//[100 2 3 4]

}
