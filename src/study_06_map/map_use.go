package main

import "fmt"

func printMap(_cityMap map[string]string)  {
  // 引用传递
  fmt.Println("====printMap====")
  for key,value := range _cityMap {
    fmt.Println("key =", key, "value =", value)
  }

  fmt.Println("====printMap end====")
}
func changeMap(_cityMap map[string]string, key string, value string)  {
fmt.Println("====changeMap====")
  // 引用传递  可以直接更改cityMap
  _cityMap[key] = value
  fmt.Println("====changeMap end====")
}
func main()  {
  cityMap := map[string]string {
    "CN":"Beijing",
    "JP":"Tokyo",
    "USA":"NewYork",
  }

  // 遍历map

  for key,value := range cityMap {
    fmt.Println("key =", key, "value =", value)
  }
  // 添加属性
  cityMap["RU"] = "Moscow"
  fmt.Println("cityMap =", cityMap)

  delete(cityMap, "RU")
  fmt.Println("cityMap =", cityMap)

  printMap(cityMap)
  changeMap(cityMap, "EN", "London")
  fmt.Println("cityMap =", cityMap)
  printMap(cityMap)

}
