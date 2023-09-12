package main

import "fmt"

func main() {
	//map类型的切片
	var a = make([]map[string]string, 3, 3) //这里将map[string]string  看做成一个整体一个类型 更好理解
	//fmt.Println(a)
	//[map[] map[] map[]]  map的默认值nil
	if a[0] == nil {
		a[0] = make(map[string]string)
		a[0]["hhhhh"] = "zhangsan"
		a[0]["66666"] = "azhe"
	}

	if a[1] == nil {
		a[1] = make(map[string]string)
		a[1]["hhhhh"] = "lisi"
		a[1]["66666"] = "dierge"
	}
	//fmt.Println(a) //[map[66666:azhe hhhhh:zhangsan] map[66666:dierge hhhhh:lisi] map[]]   最后一个是空map
	for _, v := range a {
		for key, value := range v {
			fmt.Println(key, value)
		}
	}
}
