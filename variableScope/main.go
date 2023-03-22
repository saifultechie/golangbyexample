package main

import "fmt"

var bbb string = "cricket"

func main() {
	var aaa = "saiful"
	fmt.Println(aaa)
	test()

	u2 := &subject{name: "cse"}
	fmt.Println(u2)

	// variable reassigned

	a, b := 1, 2

	b, c := 6, 7

	fmt.Println(a, b, c)
}

type subject struct {
	name string
}

func test() {
	// undefined because aaa is local scoped or function level
	// fmt.Println(aaa)

	// but bbb variable can be access because it is package level or global
	fmt.Println(bbb)
}
