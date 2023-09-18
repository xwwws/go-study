package lib1

import "fmt"


// 当前lib1包提供的api
func Lib1Test()  {
fmt.Println("Lib1Test().....")
}


// 首字母小写不会被外部访问  private
func ppLib1Test()  {
fmt.Println("ddLib1Test().....")
}


func init()  {
  fmt.Println("lib1 init...")
}
