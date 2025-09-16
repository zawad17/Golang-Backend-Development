package main
import "fmt"



func add(a , b){
	c := a + b 
}

func main() {
	//invoked a function
	add(3,4)

	func( a int, b int) {    // anoonymous func
		c := a + b
		fmt.Println(c)
	}(4,5)  //Immediate call 
 // expression  means protita line

}

a := 10 // expression
func init(){
	if a > 0{ // if  block
		fmt.Println(a)
	}
	fmt.Println("This function call first ")
}





