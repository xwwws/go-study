package main

import (
  "fmt"
)
func printArray(arr [10]int)  {
  fmt.Println("arr", arr)
  for index,value := range arr {
    fmt.Println("index =", index,"value =", value)
  }
}

func main()  {
  // 固定长度数组
  var myArray1 [10]int

  myArray2 := [10]int{1,2,3,4,5,6,7,8,9,0}

  myArray3 := [3]int{1,2,3}
  for i:= 0; i< len(myArray1); i++ {
    fmt.Println(myArray1[i])
  }
  for index,value := range myArray2 {
    fmt.Println("index =", index, "value =",value)
  }
  //  拷贝
  printArray(myArray1)
  printArray(myArray2)
  // printArray(myArray3)


  fmt.Printf("myArray1 type = %T\n",myArray1)
  fmt.Printf("myArray2 type = %T\n",myArray2)
  fmt.Printf("myArray3 type = %T\n",myArray3)


















}
