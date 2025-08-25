package main
import "fmt"

func main() {

	// < , > , <= , >= , == 
	age := 10
	if age >18 {
		fmt.Println ("You are eliguble for married")
	}else if age  == 18 {
		fmt.Println("You are almost eliguble for married")

	}else {
		fmt.Println("You are not eliguble for married")
	}

	gender := "male"
	if age == 18 || gender == "male"{
		fmt.Println("You are eliguble for married")
	}else if age > 18 || gender == "male"{
		fmt.Println("You are eliguble for married")
	}else {
		fmt.Println("You are not eliguble for married")
	}
}