package main

import (
	"fmt"
	"golangbyexample/struct/model"
)

func main() {
	p1 := model.Person{Name: "saiful", Age: 23}
	fmt.Println("the person is : ", p1)

	c1 := company{name: "bs23"}
	fmt.Println("the company ", c1)
}
