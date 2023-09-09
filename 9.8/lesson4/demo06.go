package main

import "fmt"

func main() {
	//引用传递
	//和值传递正好不一样
	//值传递是拷贝
	//引用传递是地址空间
	//因此，s1 s2 指向同一个空间

	//切片 slice 可以扩容的数组
	//但是，数组是有固定大小的，函数传递的时候要传入数组大小的
	//切片的定义与数组很类似
	s1 := []int{1, 2, 3, 4}
	fmt.Println("默认的数据", s1)
	upDate2(s1)
	fmt.Println("调用后的数据", s1)

}
func upDate2(s2 []int) {
	fmt.Println("传递的数据", s2)
	s2[0] = 100
	fmt.Println("s2修改后的数据", s2)

}
