package main

import (
  "fmt"
)

func printArray(arr []int)  {
  for index,value := range arr {
    fmt.Println("index =", index, "value =", value)
    arr[index] = 100
  }
}

func main()  {
  // 动态数组  引用传递
  myArray := []int{1,2,4,5,56,3,4,2} // 动态数组  slice

  fmt.Printf("myarray type = %T\n", myArray)  // myarray type = []int

  fmt.Println(myArray)
  fmt.Println("=====")
  printArray(myArray)
  fmt.Println("====")
  printArray(myArray)

}
