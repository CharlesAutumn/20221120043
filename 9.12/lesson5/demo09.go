package main

func main() {
	//可以自定义类型
	//type 关键字  可以定义函数、定义一个自己的变量类型
	//type出来的  打印出来的类型是 main.myInt
	//即使指定myInt 为int类型，但是还是不一样的  计算时还是需要进行类型转换

	//具体不写了
	//快速上手为主

}

func add(x, y int) int {
	return x + y
	//类型
	//函数作为另一个函数的参数
	//注意匿名函数的使用
	//定义的话，正常函数没有函数名称
	//普通的函数是无法一边定义一边在main函数中使用的，这时就要使用到匿名函数
	//定义完匿名函数  在main中直接定义完后  在{}后加上（）  ()也可以直接加入参数  称为匿名自执行函数
	//也可以var fn = func(x,y int) int{}   调用的时候使用fn(1,2)
}
