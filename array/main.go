package main

import "fmt"

func main() {
	fmt.Println("this is array section")

	// array with number of elements and actual elements

	arr := [3]int{1, 2, 3}

	fmt.Printf("the array length = %d, value is %v\n", len(arr), arr)

	// array with only actual elements

	arr2 := [...]int{1, 2, 3}

	fmt.Printf("the array length = %d, value is %v\n", len(arr2), arr2)

	// array with only numbers of elements

	arr3 := [3]int{}

	fmt.Printf("the array length = %d, value is %v\n", len(arr3), arr3)

	// without both

	arr4 := [...]int{}
	fmt.Printf("the array length = %d, value is %v\n", len(arr4), arr4)

	// arrays are value in go

	sample1 := [2]string{"a", "b"}

	fmt.Printf("sample1 before assignment : %v\n", sample1)

	sample2 := sample1
	sample2[0] = "c"
	fmt.Printf("sample2 is: %v\n", sample2)

	fmt.Printf("sample1 after assignment : %v\n", sample1)

	testFunc(sample1)

	fmt.Printf("sample1 after testing : %v", sample1)

	// different ways of array iterating

	myArr := [3]string{"saiful", "rafi", "tanmoy"}

	for i := 0; i < len(myArr); i++ {
		fmt.Printf("the value is : %s\n", myArr[i])
	}

	for _, val := range myArr {
		fmt.Printf("the value is = %s\n", val)
	}

	// 2 dimentional array

	samples := [2][3]int{{1, 2, 3}, {4, 5, 6}}

	for _, i := range samples {
		for _, j := range i {
			fmt.Printf("i = %d, j = %d\n", i, j)
		}
	}

}

func testFunc(sample [2]string) {
	sample[0] = "d"
	fmt.Printf("the test copy is : %v\n", sample)
}
