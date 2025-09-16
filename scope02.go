package main
import "fmt"

var a = 10;
var b = 20;


func main(){
	addition();
	// p:= 20;
	// if p >= 20{
	// 	q := 20;
	// 	fmt.Println("I am ",p ,"years old");
	// 	fmt.Println("I have ",q ,"years of experience");
	// }

	
	// undefined: q -> if block sesh hole close hoye ay er jonno q er value o tar position harai fele.
	/*
	fmt.Println("I have ",q ,"years of experience");
	*/
}




func addition (){
	x := 30;
	if x != 30{
		y:= 40;
		fmt.Println("Value is ",x+y)

	}
	//fmt.Println("Value is ",x+y)  //-->Extra y in expression that why sow argument error


}

//  Go te function jekhanei thakuk na seita memory run er jonno problem hobe na 
