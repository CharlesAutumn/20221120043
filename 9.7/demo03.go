package main

import "fmt"

func main() {
	var name string = "kuangshen"
	name = "zhangsan"
	var age int = 18
	fmt.Println("还会不会报错")
	fmt.Println(name)
	fmt.Println(age)
	//多变量一起定义
	//定义变量的规则还是像Java一样使用驼峰命名
	var (
		nameA string
		ageA  int
		addA  string
	)

	nameA = "kuangshen"
	ageA = 25
	addA = "zhongguo"
	fmt.Println(nameA, ageA, addA)
	//输出结果为 空格 0 空格
	//int 的默认值为0 而
	//string 的默认值为空->空格

}
