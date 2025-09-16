package main
import "fmt"

/*

Internal Memory
---------------
--> Code Segment
--> Data Segment
--> Stack
--> Heap

*/


var a = 10  // data segment part 
func add (x int, b int){      // code segment part 
	c := x + b
	fmt.Println(c)
}

func main(){    // code segment part 
	add (3,5)
	add (a,5)

}

func init(){
	fmt.Println("Print First World")
}

// if work done pop stack frame 
