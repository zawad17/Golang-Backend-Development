package main
import "fmt"

func main(){

	// add (3,4)    --> undefined error asbe 
	add := func(a int , b int){
		c := a + b
		fmt.Println(c)
	}

	add(2, 3)
}


func init(){
	fmt.Println("This line run 1st-> ")
}