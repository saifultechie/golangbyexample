package main

import (
	"fmt"
	h "golangbyexample/calculator/helper"
	n "golangbyexample/calculator/userInput"
)

func main() {
	var arType int

	fmt.Println("please select : ")
	fmt.Println("1.Addition")
	fmt.Println("2.Subtraction")
	fmt.Println("3.Multiplication")
	fmt.Println("4.Division")

	fmt.Scanf("%v\n", &arType)

	num1, num2 := n.EnterNumber()

	res := h.Cal(num1, num2, &arType)

	fmt.Println("total results is : ", res)

}
