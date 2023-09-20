package main

import "fmt"


func main()  {
  // 声明myMap1是一个map类型 key是string  value是string
  var myMap1 map[string]string
  if myMap1 == nil {
    fmt.Println("myMap1 是一个空map")
  }
  myMap1 = make(map[string]string, 3)
  myMap1["one"] = "java"
  myMap1["two"] = "c++"
  myMap1["three"] = "python"
  myMap1["four"] = "javascript"
  fmt.Println(myMap1)
  // fmt.Printf("myMap1 = %m,cap = %c\n", myMap1, cap(myMap1))

  // 声明方式2
  myMap2 := make(map[int]string)
  myMap2[1] = "java"
  myMap2[2] = "c++"
  myMap2[3] = "python"
  myMap2[4] = "javascript"
  fmt.Println(myMap2)



  // 声明方式3
  myMap3 := map[string]string{
    "key1": "value1",
    "key2": "value2",
    "key3": "value3",
    "key4": "value4",
    "key5": "value5",
  }

  fmt.Println(myMap3)
  myMap3["key6"] = "value6"
  fmt.Println(myMap3)
}
