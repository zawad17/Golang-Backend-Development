package main

import "fmt"

func sum(number1 int, number2 int) {
	result := number1 + number2
	fmt.Println("The sum is : ", result)
}

func getNumbers(num1 int, num2 int) (int, int) {
	sum := num1 + num2
	multiple := num1 * num2
	return sum, multiple

}

func main04() {
	sum(12, 34)
	sum_result, multiple_result := getNumbers(12, 34)
	fmt.Println("The sum and multiple  of two numbers is : ", sum_result, ", ", multiple_result)

}
