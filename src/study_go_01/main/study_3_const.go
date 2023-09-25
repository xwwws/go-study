package main
import "fmt"
func main () {
	/*
	iota 只能在const 中出现
	*/
	const length int  = 123
	fmt.Println(length)

	const (
		BEIJING = 10*iota
		SHANGHAI
		SHENZHEN
		GUANGZHOU
	)

	fmt.Println(BEIJING)
	fmt.Println(SHANGHAI)
	fmt.Println(SHENZHEN)
	fmt.Println(GUANGZHOU)
}