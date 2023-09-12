package main

func main() {

	//map 引用数据类型  必须初始化才能够使用
	//使用make创建map类型数据
	//var a = make(map[string]string)
	//a["hhhh"] = "hhhh"
	//fmt.Println(a)  //列出所有

	//map也支持声明的时候填充元素
	//var a = map[string]string{
	//	"你好":   "hh",
	//	"hhhh": "kkk", //注意,号
	//}
	//fmt.Println(a)

	//a:=map[string]string{
	//	"你好":   "hh",
	//	"hhhh": "kkk", //注意,号
	//}
	//fmt.Println(a)

	//v,ok := a["你好“]   查找有没有这个键值 有的话，v为值  ok为true
	//查找不到这个键值 v为默认的对应数据默认值 OK为false

	//通过delete（）删除
	//delete（a，“你好”）
}
