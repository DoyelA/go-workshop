//printf- print formatted for format specifier. %v is used to specify
//%T prints the type of variable
package main
import "fmt"
func main(){
	var name string
	name="JIIT"

	fmt.Printf("I am in %v\n", name)

	var(
		i int=20
		f float64=3.14
	)

	fmt.Printf("i: %v %T", i, i)
	fmt.Printf("f: %v %T", f, f)
}