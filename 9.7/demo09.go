package main

import "fmt"

func main() {
	const (
		a = iota
		b
		c
		d = "hhahahah"
		e
		f = 100
		g
		h = iota
		i
		//很奇妙
	)
	const (
		j = iota
		k
	)

	//观察就完了
	//开发的时候用的并不多
	fmt.Println(a, b, c, d, e, f, g, h, i)
	fmt.Println(j, k)
}
