package main

import "fmt"

func main() {
	a := 1
	b := 2
	c := 3

	if a > b && c > a {
		fmt.Println("a > b && c > a")
	}

	switch a {
	case 1:
		fmt.Println("a == 1")
		break
	case 2:
		fmt.Println("a == 2")
		break
	default:
		fmt.Println("default")

	}

	if a > b {
		fmt.Println(a)
	} else {
		fmt.Println(b)
	}
}
