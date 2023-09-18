package main
import (
  "fmt"
  "go-study/src/gocode/study_go_02/lib1"
  "go-study/src/gocode/study_go_02/lib2"
)

// 主包的init
func init()  {
fmt.Println("main init...")
}
func main()  {
  lib1.Lib1Test()
  lib2.Lib2Test()
}
