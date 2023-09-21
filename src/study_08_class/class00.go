package main

import (
  "fmt"
)
/*
  类声明
*/
// 如果类名的首字母大写,表示其他包也能范文
type Hero struct {
  // 如果类的属性首字母大写 表示该属性石对外能够访问的,否则只能够类的内部访问
  Name string
  Ad int
  Level int
}
/*
func (this Hero) Show()  {
  fmt.Println("hero =", this)
}

func (this Hero) GetName() string {
  return this.Name
}

func (this Hero) SetName(newName string)  {
  // this 是调用该方法对象的一个拷贝
  this.Name = newName
}
*/
func (this *Hero) Show()  {
  fmt.Println("hero =", this)
}

func (this *Hero) GetName() string {
  return this.Name
}

func (this *Hero) SetName(newName string)  {
  // this 是调用该方法对象的一个拷贝
  this.Name = newName
}


func main()  {
  hero := Hero{Name: "zhang3", Ad: 12, Level: 1}
  hero.SetName("li4")
  hero.Show()
  fmt.Println(hero.GetName())
  fmt.Println(hero.Ad)

}
