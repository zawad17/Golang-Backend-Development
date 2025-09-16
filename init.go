package main 
import "fmt"
 /*
func main(){
	// init()  --> can't write this
	fmt.Println("My wife name is boltu Coltu")
}

//init function --> don't need to call this function
func init(){
	fmt.Println("My Girlfriend name is Ritu")  // execute first 
}

*/

// a := 10  --> can not use in global , only use function block
var a = 10 
func main(){
	fmt.Println("Number : ",a)

}

func init(){
	fmt.Println("Number : ",a)
	a = 22
}