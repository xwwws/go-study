package main
import (
  "fmt"
  "go-study/src/study_go_02/lib1"
  "go-study/src/study_go_02/lib2"
)

// 导包方式-匿名  _ "go-study/src/gocode/study_go_02/lib1"
// 导包方式-别名  myLib1 "go-study/src/gocode/study_go_02/lib1"
// 导包方式-全量导入  . "go-study/src/gocode/study_go_02/lib1"

// 主包的init
func init()  {
fmt.Println("main init...")
}
func main()  {
  lib1.Lib1Test()
  lib2.Lib2Test()
}
