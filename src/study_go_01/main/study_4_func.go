package main
import "fmt"


func foo1(a string, b int) int {
	fmt.Println(a)
	fmt.Println(b)
	c:= 44
	return c
}
func foo2() (r1 int,r2 string) {
	fmt.Println("----foo2----")
	fmt.Println(r1)
	fmt.Println(r2)
	r1 = 1
	r2 = "adf"
	return 
}
func main () {
	c := foo1("sdf", 12)
	fmt.Println(c) 
	rd1,rd2 := foo2()
	fmt.Println(rd1,rd2)
	
	
}