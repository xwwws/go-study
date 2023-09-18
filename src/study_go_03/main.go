package main
import (
  "fmt"
)

// countAddress = &count  获取count的地址
// c = *countAddress  获取countAddress的 引用

func changeValue1(n int) {
  n = n + 1
  fmt.Println("changeValue1--- n =", n)
}

// 引用传递
func changeValue2(n *int)  {
  *n = *n + 2
  fmt.Println("changeValue2 --- n = ", n)
}

func showAddress(count *int) {
  fmt.Println("count address",count)
  fmt.Println("count",*count)
  num := 100
  fmt.Println("=====num=====")
  fmt.Println("num", num)
  fmt.Println("&num", &num)
  fmt.Println("*&num", *&num)
  fmt.Println("=====num=====")
}


func main()  {
  count := 1
  changeValue1(count)
  fmt.Println("main changeValue1 log count  = ", count)
  changeValue2(&count)
  fmt.Println("main changeValue2 log count =", count)

  fmt.Println("count",count)
  fmt.Println("&count",&count)
  showAddress(&count)
}
