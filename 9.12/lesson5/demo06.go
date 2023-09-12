package main

func main() {

	//sort!!!好用的排序包

	//升序排序，   注意可以对数、也可以对字符串
	// sort.Ints()   （）中传入切片
	//sort.Float64s()
	//sort.Strings()   基本用不到！！！了解即可
	//以上三个函数都是默认从小到大的，对字符串的排序，想想abcdefg

	//降序排序
	//sort.Sort(sort.Reverse(sort.IntSlice(intList))
	//首先，sort.IntSlice()
	//`sort.IntSlice`是Go语言中的一个类型，它是`[]int`切片类型的别名，并实现了`sort.Interface`接口，可以方便地对整数切片进行排序。
	//具体解释如下：
	//- `sort.IntSlice`是一个类型，它将`[]int`切片类型命名为`sort.IntSlice`。
	//- `[]int`表示一个整数切片，它是一个可以动态增长和缩减的可变长度的整数数组。
	//- `sort.IntSlice`实现了`sort.Interface`接口，该接口定义了排序所需的方法，包括`Len()`、`Less(i, j int) bool`和`Swap(i, j int)`。
	//使用`sort.IntSlice`类型的好处是可以直接调用`sort`包中的排序函数，而无需自己实现排序相关的方法。`sort.IntSlice`已经为我们封装了这些方法，使得对整数切片的排序变得非常简洁和方便。
	//package main
	//
	//import (
	//	"fmt"
	//	"sort"
	//)
	//
	//func main() {
	//	nums := []int{4, 2, 7, 1, 5}
	//	sort.Sort(sort.IntSlice(nums))
	//	fmt.Println(nums) // 输出：[1 2 4 5 7]
	//}

}
