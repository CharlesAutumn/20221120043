package main

import "fmt"

func main() {

	fmt.Println(add(1, 2))

}

func add(a, b int) int { //结构有点特殊，要记忆一下
	c := a + b
	return c

}
