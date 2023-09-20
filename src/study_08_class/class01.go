package main

import (
  "fmt"
)
/*
  类继承
*/
type Human struct {
  name string
  gender string

}
func (this *Human) Eat()  {
  fmt.Println(this.name,"Eat...")
}
func (this *Human) Walk()  {
  fmt.Println(this.name,"Walk...")
}


type SuperMan struct {
  Human  // SuperMan 类继承了Human类的属性和方法
  level int
}

// 重定义父类的方法Walk()
func (this *SuperMan) Walk()  {

    fmt.Println(this.name,"SuperMan.Walk...")
}

// 子类的新方法
func (this *SuperMan) Fly()  {

      fmt.Println(this.name,"SuperMan.Fly...")
}

func (this *SuperMan) GetLevel() int {
  return this.level
}

func (this *SuperMan) Show() {
  fmt.Println("====SuperMan====")
  fmt.Println("name = ", this.name)
  fmt.Println("gender = ", this.gender)
  fmt.Println("level = ", this.level)
  fmt.Println("====SuperMan end====")
}

func main()  {

  h:= Human{"zhang3","male"}
  h.Eat()
  h.Walk()

  supMan := SuperMan{Human{"Sli4","Smale"},100}
  supMan.Walk()
  supMan.Eat()
  supMan.Fly()
  supMan.Show()

}
