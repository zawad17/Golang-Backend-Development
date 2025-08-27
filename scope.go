package main
import "fmt"

var a = 10;
var b = 20;

func add (x int, y int){
	z := x + y;
	fmt.Println("Sum is : ",z)
}

func main(){
	p := 30;
	q := 40;


	add(p, q);
	add(p, a);
	add(a, b);
	add(q, b);
	add(a, z);
}



