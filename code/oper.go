package  main
import "fmt"
func main() {
	var a int = 5
	var b int = 10
	var c float32 = 23.12
	x,y :=14,15
	const pi float32 = 3.13
	isbool :=true
	hate :=false
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(x,",",y)
	fmt.Println(pi)

	fmt.Println("a+b=", a+b)
	fmt.Println("b-c=", b-x)
	fmt.Println("x*y=", x*y)
	fmt.Println("b/a=", b/a)
	fmt.Println("area of circle=", 2*pi*c*c)

	fmt.Println(isbool && hate )
	fmt.Println(isbool || hate)
	fmt.Println(!isbool)
	fmt.Println(!hate)
}