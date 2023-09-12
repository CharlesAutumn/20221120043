package main

import "fmt"

func main() {

	var arr3 [3]int
	arr3[0] = 1
	arr3[1] = 2
	arr3[2] = 3

	var arr1 = [4]int{1, 2, 3, 4}

	arr2 := [3]int{1, 2, 3}

	var arr4 = [...]int{12, 2345}
	fmt.Println(arr4)

	arr6 := [...]int{0: 1, 5: 2}
	fmt.Println(arr6)

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)

	//var str string = strconv.Itoa(x)   整形转化为字符串
	//对应的  golang strconv.ParseInt 是将字符串转换为数字的函数,参数1 数字的字符串形式,参数2 数字字符串的进制
	//`strconv.ParseInt()`函数是Go语言中用于将字符串转换为整数类型的函数。
	//具体解释如下：
	//- `strconv`是Go语言的标准库中提供字符串和基本数据类型之间相互转换的包。
	//- `ParseInt()`是`strconv`包中的一个函数，用于将字符串解析为整数类型。
	//- `ParseInt()`函数的签名如下：`func ParseInt(s string, base int, bitSize int) (i int64, err error)`。
	//- `s`是要解析的字符串。
	//- `base`指定进制，可以是2到36之间的任意整数，或者0表示根据字符串前缀来判断进制。通常使用10表示十进制。
	//- `bitSize`指定结果的位大小，可以是0、8、16、32、64中的一个。0表示根据字符串自动选择位大小。
	//- `i`是解析后的整数值。
	//- `err`是解析过程中可能遇到的错误。

	//在Go语言中，`make()`函数用于创建切片、映射和通道等引用类型的对象。
	//具体解释如下：
	//- `make()`函数的语法是：`make(T, length, capacity)`。
	//- `T`表示要创建的对象的类型，可以是切片、映射或通道等引用类型。
	//- `length`表示切片的初始长度或映射的初始容量。
	//- `capacity`只适用于切片，表示切片的预分配容量。

	//在Go语言中，`copy()`函数用于将源切片的元素复制到目标切片中。
	//具体解释如下：
	//- `copy()`函数的语法是：`copy(dst []T, src []T) int`。
	//- `dst`表示目标切片，即要将元素复制到的切片。
	//- `src`表示源切片，即要从中复制元素的切片。
	//- `T`表示切片中元素的类型。
	//- `copy()`函数返回实际复制的元素个数。

	//在Go语言中，`cap()`函数用于获取切片的容量。
	//具体解释如下：
	//- `cap()`函数的语法是：`cap(s []T) int`。
	//- `s`表示要获取容量的切片。
	//- `T`表示切片中元素的类型。
	//- `cap()`函数返回切片的容量。

	//知识点：
	//s|n|    切片s中索引位置为n的项  单个
	//s|:|    从切片s的索引位置0到len(s)-1 处所获得的切片  所有
	//s|low:|   从切片s的索引位置 low 到len(s)-1 处所获得的切片  包前不包后
	//s|:high| 从切片s的索引位置 0到high 处所获得的切片，len=high   也是包前不包后  high是索引
	//s|low: high|  从切片s的素引位置 Iow 到high 处所获得的切片，len-high-low
	//s|low: high:max|  从切片s的素引位置 low 到high 处所获得的切片，len-high-low, cap=max-low

	//在Go语言中，`append()`函数用于将一个或多个元素追加到切片的末尾。当需要将一个切片的所有元素追加到另一个切片时，可以使用`...`操作符将切片展开为单个元素序列。
	//具体解释如下：
	//- `append()`函数的语法是：`s1 = append(s1, s2...)`。
	//- `s1`是目标切片，即要将元素追加到的切片。
	//- `s2`是源切片，即要追加的切片。
	//- `...`操作符用于展开切片，将切片中的所有元素作为单个元素序列进行追加。

	//在Go语言中，`sort.Ints()`函数用于对整型切片进行升序排序。
	//具体解释如下：
	//- `sort.Ints()`函数的语法是：`sort.Ints(slice)`。
	//- `slice`是待排序的整型切片。

	//在Go语言中，`map`是一种无序的键值对集合。它类似于其他编程语言中的哈希表或字典。
	//具体解释如下：
	//- `map`的声明语法是：`var m map[keyType]valueType`。
	//- `m`是一个`map`类型的变量。
	//- `keyType`是键的类型，可以是任意可比较类型（例如整数、字符串、结构体等）。
	//- `valueType`是值的类型，可以是任意类型。
	//m := make(map[string]int)
	//
	//	m["apple"] = 5
	//	m["banana"] = 3
	//
	//	fmt.Println(m) // 输出：map[apple:5 banana:3]

	//在Go语言中，可以使用类型转换将字符串转换为字节数组。
	//具体解释如下：
	//- 假设`s`是一个字符串，`arr`是一个字节数组。
	//- 使用类型转换的语法是：`arr := ([]byte)(s)`。   不是转换成切片
	//- `(s)`将字符串`s`转换为字节数组类型`[]byte`。
	//- `[]byte`是一个字节数组类型，用于存储原始数据的字节表示。

	//这个技术不错
	//func character(s string) byte {
	//	// write code here
	//	arr := ([]byte)(s)
	//	count := 0
	//	var ans byte
	//	charaMap := make(map[byte]int)
	//	for _, j := range arr {
	//
	//		charaMap[j]++
	//		if charaMap[j] > count {
	//			count = charaMap[j]
	//			ans = j
	//		}
	//	}
	//
	//	return ans
	//}

	//在Go语言中，可以使用以下方式判断一个`map`中是否存在指定的`key`：
	//具体解释如下：
	//- `m`是一个`map`类型的变量。
	//- `key`是要检查的键。
	//- `_`是一个特殊的占位符，用于忽略不需要的值。
	//- `ok`是一个布尔值，用于表示`key`是否存在于`map`中。

}
