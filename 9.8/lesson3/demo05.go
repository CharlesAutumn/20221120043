package main

import "fmt"

func main() {

	//穿透!!!!!很好玩，正常Java，c是自动穿透，
	//但是go是需要用fallthrough,感觉是真的体现了简化操作

	//同时，可以在fallthrough的路径上，使用break来停止穿透
	a := false
	switch a {
	case false:
		fmt.Println("case的条件是false")
		fallthrough //！！！！！只是会穿透到下一个case，不是一直到最后
	case true:
		if a == false {
			break
		}
		fmt.Println("case条件是true")

	}

}
