package main

import (
  "fmt"
)
func main()  {
  // make(type,len(长度),cap(容量))
  //
  /*
    append(to, 0) 向to最后追加 一个元素  元素内容为0
    每次给silce append的时候 在slice的最后追加一个元素, 如果容量(cap)不够  继续向后追加 cap个容量 以此类推
  */
  var numbers = make([]int, 3, 5)

  fmt.Printf("len = %d, cap = %d slice = %v\n", len(numbers),cap(numbers), numbers)
  // 向numbers切片追加一个元素1 numbers
  numbers = append(numbers, 1)
  fmt.Printf("len = %d, cap = %d slice = %v\n", len(numbers),cap(numbers), numbers)
  //len = 3, cap = 5 slice = [0 0 0]


  numbers = append(numbers, 2)
  fmt.Printf("len = %d, cap = %d slice = %v\n", len(numbers),cap(numbers), numbers)
  // len = 4, cap = 5 slice = [0 0 0 1 2]


  numbers = append(numbers, 3)
  fmt.Printf("len = %d, cap = %d slice = %v\n", len(numbers),cap(numbers), numbers)
  // len = 6, cap = 10 slice = [0 0 0 1 2 3]


  // make(type, len)  当不使用cap时   cap 与 len 相同
  var numbers2 = make([]int, 3)


  fmt.Printf("len = %d, cap = %d slice = %v\n", len(numbers2),cap(numbers2), numbers2)
  // len = 3, cap = 3 slice = [0 0 0]

  numbers2 = append(numbers2, 1)
  fmt.Printf("len = %d, cap = %d slice = %v\n", len(numbers2),cap(numbers2), numbers2)
  // len = 4, cap = 6 slice = [0 0 0 1]

}
