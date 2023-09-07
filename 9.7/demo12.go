package main

import "fmt"

func main() {

	var str string
	str = "Hello,kuangshen"
	fmt.Printf("%T,%s\n", str, str)

	//单引号
	v1 := 'A'
	v2 := "A"

	//编码表 ACSSI
	fmt.Printf("%T,%c\n", v1, v1) //int32  65   但是默认int是64的   %d也是可以的，输出为65
	fmt.Printf("%T,%s\n", v2, v2)

	//对于v3 := '中'  也可以打印出来，对应的%d是有的，int32
	//原因是即使ACSSI中没有 中 这个字，但是在另一个编码表中存在这个
	//所有中国字的编码表：GBK
	//全世界的编码表：Unicode编码表    go 支持！！！

	//转义字符
	//  跟c一样，   \    ,这个符号

}
