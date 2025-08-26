package main
import "fmt"
func getName()string{
	var name string 
	fmt.Println("Enter your name:")
	fmt.Scanln(&name)
	return name 
}

func getNumber()(int, int){
	var num1, num2 int
	fmt.Println("Enter two numbers:")
	fmt.Scanln(&num1, &num2)
	return num1, num2
}

func welcome(name string){
	fmt.Println("Welcome", name)
}

func add(num1, num2 int)(int){
	return num1 + num2
}
func display(name string, sum int, num1 int, num2 int){
	fmt.Println("Hello", name)
	fmt.Println("The sum of", num1, "and", num2, "is", sum)
}

func main(){
	name :=getName()
	num1, num2 := getNumber()
	welcome(name)
	sum := add(num1,num2)
	display(name, sum, num1, num2)
}