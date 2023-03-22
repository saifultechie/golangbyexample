package userinput

import "fmt"

func EnterNumber() (int, int) {
	var num1, num2 int

	fmt.Print("please enter two number :")
	fmt.Scanf("%v", &num1)
	fmt.Scanf("%v", &num2)

	return num1, num2
}
