
package main

import (
  "fmt"
)
func main()  {
  s := []int{1,2,3,4,5} // len = 5 cap = 5


  // 截取  s[start : end]  start:开始下标  end:结束下标
  // 可以单独只传入一个参数
  // s[:end] 从下标0开始 截取到下标为end的元素  包括end
  // s[start:] 截取从start开始到末尾 不包括start


  // 截取的是每个元素的引用地址
  // 每次更改某个元素  都会影响到所有该元素的引用



  // copy 可以将底层数组的slice一起进行拷贝



  s1:= s[0 : 4]
  fmt.Println("s1 =", s1)
  // s1 = [1 2 3 4]


  s2:= s[1:4]
  fmt.Println("s2 =", s2)
  // s2 = [2 3 4]


  s3:= s[:4]
  fmt.Println("s3 =", s3)
  // s3 = [1 2 3,4]


  s4:= s[3:]
  fmt.Println("s4 =", s4)
  // s1 = [4 5]
  fmt.Println("=====change=====")
  s[3] = 100

  fmt.Println("s =", s)
  // s = [1 2 3 100 5]
  fmt.Println("s1 =", s1)
  // s1 = [1 2 3 100]
  fmt.Println("s2 =", s2)
  // s2 = [2 3 100]
  fmt.Println("s3 =", s3)
  // s3 = [1 2 3 100]
  fmt.Println("s4 =", s4)
  // s4 = [100 5]



  fmt.Println("=====copy=====")


  /*
    copy(to, form) 深拷贝  将from 拷贝到to
  */

  slice := make([]int, 5) // [0 0 0 0 0]

  copy(slice, s)
  fmt.Println("slice = ", slice)
  fmt.Println("s = ", s)

}
