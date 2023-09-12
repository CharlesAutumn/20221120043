package main

func main() {

	//声明切片4种方式
	//var arr1 []int
	//var arr2 = []int{1,2,3,4}
	//arr3 :=[]string{"hhh","hhhhhh"}
	//var arr4 = []int{1:3,4:5}

	//make()

	//var sliceA = make([]int, 4, 8)
	//这个时候就可以对sliceA[1]进行赋值操作

	//var sliceB []int    这个时候长度、容量都是0
	//需要使用append(sliceB,12)  append(sliceB,13)
	//append(sliceB,12,14,25)

	//append(sliceA,sliceB...)合并切片

	//切片的扩容策略
	//var sliceA []int
	//for i:=0;i<10;i++{
	// sliceA = append(sliceA,i)
	//}
	//在容量是不够的时候，直接扩容操作，x2

	//copy()

	//切片是引用型类型  对副本操作，会改变原来的值  特性

	//如果使用copy() 就不会发生以上情况
	//sliceA := []int{1,2,3,4}
	//sliceB := make([]int ,4,4)
	//copy(sliceB,sliceA)  //注意谁在前

	//在go中没有特定的删除切片元素的方法，但是可以利用特性append
	// a:=[]int{1,2,3,4,5,6,7}
	//a = append(a[:2],a[3:]...)  删除索引为2的元素

	//修改字符串中的字符
	//go语言中字符串时无法修改的，这个时候是可以将字符串转化成切片，对其进行修改
	//两种  一种转化成byte切片  没有汉字就用byte  有汉字就用另一种  rune
	//s1 :="big"
	//byteStr := []byte(s1)
	//byteStr[0] = 'p'
	//fpl(string(byteStr))
	//
	//s2:="你好golang"
	//runeStr := []rune(s2)
	//runeStr[0] = '大'
	//fpl(string(runeStr))
	//
	//其实以上没有必要  毕竟可以对s1、s2直接重新赋值
}
