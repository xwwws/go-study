package main

import (
  "fmt"
)
func main()  {
  // slice 声明方式
  // 声明方式1
  // 声明slice1是个切片,并且初始化,默认是1,2,3 长度len是3
  slice1 := []int{1,2,3}
  fmt.Printf("len = %d,slice = %v\n", len(slice1),slice1)

  // 声明方式2

  // 声明slice2是个切片,并没有给slice2分配内存
  var slice2 []int
  // slice2[0] = 1 // 报错

  // 给slice开辟3个空间 内容是0
  slice2 = make([]int, 3)
  slice2[0] = 1
  slice2[1] = 2
  slice2[2] = 3
  // slice2[3] = 4 // 报错


  fmt.Printf("len = %d,slice = %v\n", len(slice2),slice2)


    // 声明方式3
    // 声明slice3是个切片,且长度是4,内容是0
    var slice3 []int = make([]int, 4)

    fmt.Printf("len = %d,slice = %v\n", len(slice3),slice3)

    // 声明方式4
    // 声明slice4是个切片,且长度是4,内容是0
    slice4 := make([]int, 3)
    fmt.Printf("len = %d,slice = %v\n", len(slice4),slice4)

    // 判断slice是否为空 length为 0
    if slice1 == nil {
      fmt.Println("slice1是空切片")
    } else {
      fmt.Println("slice1是有空间的")
    }
}
