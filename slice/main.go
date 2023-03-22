package main

import (
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
)

func main() {
	// there are four ways to declare slice

	slice1 := []int{}
	fmt.Println("slice1 length is: ", len(slice1))
	fmt.Println("slice1 capacity is: ", cap(slice1))
	fmt.Println("slice1 is: ", slice1)

	letters := []int{1, 2, 3}
	letters = append(letters, 4)

	fmt.Println("the initialize element is : ", letters)
	fmt.Println("the initialize element length is : ", len(letters))
	fmt.Println("the initialize element capacity is : ", cap(letters))

	// create re-slice from array

	numbers := [5]int{1, 2, 3, 4, 5}

	// both start and end

	num1 := numbers[2:4]

	fmt.Println("the reslice is one :", num1)

	// only start

	num2 := numbers[2:]
	fmt.Println("the re-slice from start : ", num2)

	// only end

	num3 := numbers[:2]
	fmt.Println("the re-slice from end: ", num3)

	numbers[3] = 8

	num3[0] = 10

	fmt.Println("the re-slice from start : ", num3)

	fmt.Println(numbers)

	// slice create from builtin function make

	slice2 := make([]int, 3, 5)

	fmt.Println("the length and the capacity of slice", len(slice2), cap(slice2))

	// slice create from builtin function make ommited cap value

	slice3 := make([]int, 4)

	fmt.Println("the slice3 := ", slice3)
	fmt.Println("the length and the capacity of slice3 ", len(slice3), cap(slice3))

	// talk with length and capacity

	numberss := make([]int, 3, 5)

	fmt.Println("slice length is : ", len(numberss))
	// the underline code is raise runtime error
	// numberss[4] = 90

	// fmt.Println("the new slice is ", numberss)

	// if i want to reslove this then
	//Increasing the length from 3 to 5

	numberss = numberss[0:5]
	numberss[4] = 10
	fmt.Println("slice length is : ", len(numberss))
	fmt.Println("slice is : ", numberss)

	// accessing and modifying a slice

	arrr := [5]int{1, 2, 3, 4, 5}

	slices := arrr[:]

	fmt.Println("at first create a slice from array :")
	fmt.Println("the array is: ", arrr)
	fmt.Println("the slice is : ", slices)

	// modify the array

	arrr[1] = 7

	fmt.Println("modified the array :")
	fmt.Println("the array is: ", arrr)
	fmt.Println("the slice is : ", slices)

	// modify the slice

	slices[1] = 2

	fmt.Println("modified the slices :")
	fmt.Println("the array is: ", arrr)
	fmt.Println("the slice is : ", slices)

	// append single elements into slice

	single := []int{1, 2}

	single = append(single, 3)

	// append multiple element

	single = append(single, 4, 5)

	// append a slice into another slice

	number1 := []int{1, 2, 3}

	number2 := []int{4, 5}

	number1 = append(number1, number2...)

	fmt.Println(number1)

	// copy a slice in built in function copy
	// this function return number of copied elements
	// this function takes two argument (dst,src)
	// no reflect for change any src or vice varca

	sl := []int{1, 2, 3}

	sl1 := make([]int, 3)

	numberOfelements := copy(sl1, sl)

	fmt.Println("the number of copied element : ", numberOfelements)

	fmt.Println("the src slice : ", sl)

	fmt.Println("the dst slice : ", sl1)

	sl[0] = 19

	fmt.Println("after modifying src slice")
	fmt.Println("the src slice : ", sl)

	fmt.Println("the dst slice : ", sl1)

	// nill slice

	var nillSlice []int

	nillSlice = append(nillSlice, 10)

	fmt.Println(nillSlice)

	// 2D slice

	twoDsli1 := make([][]int, 2)

	for i := range twoDsli1 {
		twoDsli1[i] = make([]int, 3)
	}

	fmt.Println(twoDsli1)

	// another way

	twoSli2 := make([][]int, 2)

	twoSli2[0] = []int{1, 2, 3}
	twoSli2[1] = []int{4, 5, 6}

	fmt.Println(twoSli2)

	// check item into a slice

	subject := []string{"cse", "bangla", "ict", "english"}

	res := findElement(subject, "icts")

	if res {
		fmt.Println("the value is exist")
	} else {
		fmt.Println("the value is not find")
	}

	// find element and delete

	before := []int{1, 2, 3, 2, 4, 5, 6, 2}

	after := findAndDelete(before, 2)

	fmt.Println("after ", after)
	fmt.Println(before)

	// find delete without modify original slice

	before1 := []int{1, 2, 3, 2, 4, 5, 6, 2}

	after1 := deleteWithoutModify(before1, 2)

	fmt.Println(after1)
	fmt.Println("orginal", before1)

	// slice to json convert

	sampleSlice := []string{"saiful", "rafi", "tanmoy"}

	fmt.Println("before convert json : ", sampleSlice)

	j, err := json.Marshal(sampleSlice)

	if err != nil {
		fmt.Printf("Error : %s", err.Error())
	} else {
		fmt.Println(string(j))
	}

	// create a slice from a struct

	type employee struct {
		name string
		age  int
	}

	employess := make([]employee, 3)

	employess[0] = employee{name: "saiful", age: 25}
	employess[1] = employee{name: "rafi", age: 36}
	employess[2] = employee{name: "jasim", age: 25}

	fmt.Println(employess)

	// create slice of map

	maps := make([]map[string]string, 3)

	// create map

	map1 := make(map[string]string)
	map1["key1"] = "footbal"

	map2 := make(map[string]string)
	map2["key2"] = "cricket"

	map3 := make(map[string]string)

	map3["key3"] = "badminton"

	maps[0] = map1
	maps[1] = map2
	maps[2] = map3

	for _, item := range maps {
		fmt.Println(item)
	}

	// create slice of boolean

	// first way

	var first_slice []bool
	first_slice = append(first_slice, true)
	first_slice = append(first_slice, false)

	fmt.Println("first boolean slice :")
	fmt.Println(first_slice)

	// second way

	second_slice := make([]bool, 3)
	second_slice[0] = true
	second_slice[1] = false

	fmt.Println("second way boolean slice :")
	fmt.Println(second_slice)

	// sort slices

	elements := []int{1, 2, 4, 3, 6, 9, 5}

	sort.Slice(elements, func(i, j int) bool {
		return elements[i] < elements[j]
	})

	fmt.Println(elements)

	// sort slices from specific length

	elements1 := []int{1, 2, 4, 3, 6, 9, 5}

	sort.Slice(elements1[3:], func(i, j int) bool {
		return elements1[3+i] < elements1[3+j]
	})

	fmt.Println(elements1)

	// array of numbers convert to string

	numbArr := []int{1, 2, 3}

	output := ""

	for _, item := range numbArr {
		temp := strconv.Itoa(item)
		output += temp
	}

	fmt.Println(output)

}

func deleteWithoutModify(s []int, item int) []int {
	findArr := make([]int, len(s))
	index := 0
	for _, val := range s {
		if val != item {
			findArr[index] = val
			index++
		}
	}

	return findArr[:index]
}

func findAndDelete(items []int, search int) []int {
	index := 0
	// newSlice := make([]int, len(items))
	// 1, 2, 3, 2, 4, 5, 6, 2
	for _, val := range items {
		if val != search {
			items[index] = val
			index++

		}
	}

	fmt.Println("ists", items)

	return items[:index]
}

func findElement(elements []string, val string) bool {
	for _, item := range elements {
		if item == val {
			return true
		}
	}

	return false
}
