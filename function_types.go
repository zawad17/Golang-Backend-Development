package main
import "fmt"

// parameter and argument 
/*
func add ( a int  , b int){ // parameter  --> function declare time variable
	
	c := a + b 
	fmt.Println(c) 

}

func main (){
	add (2, 5 ) // argument  --> function call time value pass  

}


*/


func callFunction (value string, fn func(value string)string) string {    // function er return type likhte hobe + argument func er o return type likhte hobe
	return fn(value)
}


func greet (name string) string {
	return name 
}

func main(){
	x := callFunction("zawad",greet)
	fmt.Println(x)
}