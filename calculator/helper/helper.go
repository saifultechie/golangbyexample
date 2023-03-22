package helper

var myMap = map[int]string{
	1: "+",
	2: "-",
	3: "*",
	4: "/",
}

func Cal(num1 int, num2 int, cal_type *int) int {
	var total int
	switch *cal_type {
	case 1:
		total = num1 + num2
	case 2:
		total = num1 - num2

	case 3:
		total = num1 * num2
	case 4:
		total = num1 / num2
	}

	return total
}
