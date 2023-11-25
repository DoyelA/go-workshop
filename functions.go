package main
import "fmt"

type rectangle struct{
	length, breadth int
}

func calculateSum(num1, num2 int) (int){
	return num1+num2
}

func (r rectangle) area() int{
	return r.length*r.breadth
}

func main(){
	total:=calculateSum(10, 20)
	fmt.Printf("Total: %d\n", total)

	r:=rectangle{length: 10, breadth: 20}
	fmt.Printf("Area: %v\n", r.area())
}