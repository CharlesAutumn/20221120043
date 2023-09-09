package main

import "fmt"

func main() {
	//switch

	var score int = 90
	switch score {
	case 90:
		fmt.Println("A")
	case 80:
		fmt.Println("B")
	case 70:
		fmt.Println("C")
	case 40, 50, 60:
		fmt.Println("D")
	default:
		fmt.Println("F")

	}

	switch { //默认是true
	case false:
		fmt.Println("false")
	case true:
		fmt.Println("true")
	default:
		fmt.Println("其他")

	}

}
