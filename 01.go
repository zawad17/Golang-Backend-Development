package main

import "fmt"

func main() {
	fmt.Println("Hello, Zawad!")
	fmt.Println(subtract(10, 5))
	fmt.Println(add(3, 5))

}
func add(a int, b int) int {
	return a + b
}
func subtract(a int, b int) int {
	return a - b
}
