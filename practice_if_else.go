package main

import "fmt"

func mainPractice() {

	//Basic level 01 if else condition

	/*
		var number int
		fmt.Println("Enter a number:")
		fmt.Scanln(&number)
		if number % 2== 0{
			fmt.Println("The number is even")
		}else{
			fmt.Println("The number is odd")
		}
	*/

	//Basic level 02 if else condition
	var username string
	var password string
	fmt.Println("Enter your username:")
	fmt.Scanln(&username)
	fmt.Println("Enter your password: ")
	fmt.Scanln(&password)
	if username == "admin" && password == "12345abc" {
		fmt.Println("Login Successful")
	} else {
		fmt.Println("Login Failed")
	}

}
