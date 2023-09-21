package main

import (
  "fmt"
)
// defer 表示一个函数在执行完成之后之执行的方法
// 执行顺序是在函数所有内容执行完成之后  (return之后)
// 不影响return 的值

func func1() int {
  fmt.Println("f1 end")
  return 1
}
func func2() int {
  fmt.Println("f2 end")
  return 2
}
func func3() int {
  fmt.Println("f3 end")
  return 3
}


func deferFn() int {
  defer func1()
  defer func2()
  defer func3()
  fmt.Println("defferFn end")
  return 4
}


func main()  {
  defer fmt.Println("main end")
  fmt.Println("main")
  res := deferFn()
  fmt.Println(res)
}
