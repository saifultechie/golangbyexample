package main

import "fmt"

func main() {
	// if and else condition

	a := 12

	if a > 10 {
		fmt.Println("a is greater than 10")
	} else {
		fmt.Println("a is less than 10")
	}

	// if , else if and else

	age := 19

	if age < 18 {
		fmt.Println("you are young man")
	} else if age > 18 && age < 30 {
		fmt.Println("you are adult man")
	} else {
		fmt.Println("you are old man")
	}

	// if with statement and condition together

	if a := 6; a > 5 {
		fmt.Println("your age is greater then 5")
	} else if a > 7 && a < 10 {
		fmt.Println("your age is greater ")
	} else {
		fmt.Println("your age is less then 5")
	}

	// for loop

	// sample for loop

	for i := 1; i < 5; i++ {
		fmt.Printf("the iteration is : %v\n", i)
	}

	// for loop as while loop

	i := 1

	for i < 5 {
		fmt.Printf("the iteration value is: %v\n", i)
		i++
	}

	// for loop like as block

	j := 0

	for {
		fmt.Printf("the block is : %v\n", j)
		j++
		if j > 5 {
			break
		}
	}

	// nested loop

	for i := 0; i < 3; i++ {
		fmt.Printf("outer loop iteration: %d\n", i)

		for j := 0; j < 2; j++ {
			fmt.Printf("i = %d, j = %d\n", i, j)
		}

	}

	// switch statement and expression

	switch ch := 'a'; ch {
	case 'a':
		fmt.Println("this is a")

	case 'b', 'c':
		fmt.Println("b or c")
	default:
		fmt.Println("unkonown")
	}

	// omited switch statement and expression

	// when switch expression is omited then case type default boolean

	ages := 17

	switch {
	case ages < 18:
		fmt.Println("kids")

	case ages > 18 && ages < 40:
		fmt.Println("young")

	default:
		fmt.Println("old man")

	}

	// only switch statement

	switch age := 19; {
	case age < 18:
		fmt.Println("kids")

	case age > 18 && age < 40:
		fmt.Println("young")

	default:
		fmt.Println("unmatched")
	}

	// for range with array/slice

	// letters := []string{"a", "b", "c"}

	// with index and value

	// for index, val := range letters {

	// 	fmt.Printf("index = %d , value = %s\n", index, val)
	// }

	// with only index

	// for index := range letters {
	// 	fmt.Printf("index = %d\n", index)
	// }

	// only value

	// for _, value := range letters {
	// 	fmt.Printf("value = %s\n", value)
	// }

	// for range into string

	// sample := "saiful"

	// for index, cha := range sample {
	// 	fmt.Printf("index = %d, character = %s\n", index, string(cha))
	// }

	// for range with map

	location := map[string]string{
		"name": "saiful",
		"dept": "cse",
		"dist": "tangail",
	}

	for key, value := range location {
		fmt.Printf("the key is: %s, value is : %s\n", key, value)

	}

}
