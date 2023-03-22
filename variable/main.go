package main

import "fmt"

func main() {

	// single variable declaration without initial value
	// var firstName string
	// var isTrue bool
	// var students [5]int
	// fmt.Println(firstName, isTrue, students)
	// output above this program is "". when the variable only declare then the default value assigned by type
	// it means int -> 0
	// string means -> ""
	// boolean means -> false
	// array [5] means -> [0,0,0,0,0]

	// single variable declare with initial values

	// var num1 int = 4

	// var num2 = 5
	// myHobby := "football"

	// fmt.Println(num1, num2)
	// fmt.Println(myHobby)

	// multiple variable declaration without initial value

	// var aaa, bbb int

	// fmt.Println(aaa, bbb)

	// multiple variable declaration with initial value

	// var aa, bb = 40, 50

	// fmt.Println(aa, bb)

	// declare multiple variable with different type

	var (
		name string
		age  int = 26
		dept     = "cse"
	)

	fmt.Println(name, age, dept)

}
