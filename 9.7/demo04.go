package main

import "fmt"

func main() {
	//:= 叫做 短变量  自动推导  这个符号只能对初始化的变量
	//如果先定义了var name string = "kuangshen"
	//再使用 name :="kaungshen"  不好使的，报错
	name := "kuangshen" // var name string = "kuangshen"
	age := 18

	fmt.Println(name, age)
	fmt.Printf("%T,%T", name, age)
	//打印类型用%T

}
