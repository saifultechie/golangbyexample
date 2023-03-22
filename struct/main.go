package main

import "fmt"

type player struct {
	name string
	age  int
}

func main() {
	// strcut creation

	type employee struct {
		name   string
		age    int
		salary int
	}

	// create struct variable

	emp1 := employee{}

	fmt.Println("the employee one ", emp1)

	emp2 := employee{name: "saiful", age: 26, salary: 56}
	fmt.Println("the second employee is ", emp2)

	emp3 := employee{name: "rafi", age: 26}

	// access and set
	n := emp3.name
	emp3.name = "jahn"

	fmt.Println("the third employee is ", emp3, n)

	// struct to pointer

	// there are two ways to create a struct type pointer

	// 1. & operator 2. new keyword

	empP := &emp2

	fmt.Println("the first struct type pointer", empP)
	// struct pointer can also be directly created as well
	emp5 := &employee{name: "jasim", age: 21, salary: 35}

	fmt.Println("the second struct type pointer : ", emp5)

	// using new keyword

	empP2 := new(employee)
	fmt.Printf("struct pointer two using new keyword %p\n", empP2)

	// print the struct variable with different format identifier

	emp6 := employee{name: "tuhin", age: 26, salary: 40}
	fmt.Printf("the value without key : %v\n", emp6)
	fmt.Printf("the value with key : %+v\n", emp6)
	fmt.Println("the value without key ", emp6)

	// anonymous field name for struct

	// where have no name the field name will be type name

	type student struct {
		string
		Age    int
		salary int
	}

	s1 := student{string: "jakir", Age: 34, salary: 40}
	fmt.Println(s1)

	// nested struct

	type address struct {
		city     string
		postcode int
	}

	type teacher struct {
		name    string
		age     int
		cgpa    float32
		address address
	}

	ad1 := address{city: "tangail", postcode: 1976}
	t1 := teacher{name: "saiful", age: 26, cgpa: 3.73, address: ad1}
	fmt.Println("the nested obj is : ", t1.address.city)

	// struct is a value type not reference type

	p1 := player{name: "messi", age: 38}
	fmt.Printf("the player is %+v\n", p1)

	p2 := p1
	p2.name = "cr7"

	fmt.Printf("the player2 name is changed for the new assign variable : %+v\n", p2)
	fmt.Printf("the player1 name is not changed for the new assign variable : %+v\n", p1)

	fmt.Printf("this is changed for the again copy variable : not changed original value %+v\n", formatName(p2))

}

func formatName(emp player) player {
	emp.name = "neymer"
	return emp
}
