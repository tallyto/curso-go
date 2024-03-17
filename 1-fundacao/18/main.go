package main

import "fmt"

type x interface {
}

func main() {
	var x interface{} = 10
	fmt.Println(x)
	var y interface{} = "hello"
	fmt.Println(y)
	showType(x)
	showType(y)
}

func showType(t interface{}) {
	fmt.Printf("O tipo da variável é %T e o valor é %v\n", t, t)
}
