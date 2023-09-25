package main

import (
  "fmt"
)
/*
  类  多态
*/
// 本质是一个指针
type IAnimal interface {
  Sleep()
  GetColor() string // 获取动物颜色
  GetType() string //  获取动物种类
}

// 具体的类 Cat
type Cat struct {
  color string
}

func (this *Cat) Sleep()  {
  fmt.Println("Cat sleep...")
}

func (this *Cat) GetColor() string  {
  return this.color
}

func (this *Cat) GetType() string  {
  return "Cat"
}

// dog 类
type Dog struct {
  color string
}
func (this *Dog) Sleep()  {
  fmt.Println("Dog sleep...")
}

func (this *Dog) GetColor() string  {
  return this.color
}
func (this * Dog) GetType() string {
  return "Dog"
}



func showAnimal (animal IAnimal)  {
  animal.Sleep()
  fmt.Println("color = ",animal.GetColor())
  fmt.Println("type = ",animal.GetType())
}


func main()  {
  /*
    var animal1 IAnimal
    animal1 = &Cat{"red"}
    animal1.Sleep() // 调用的是cat的sleep
    var animal2 IAnimal
    animal2 = &Dog{"yellow"}
    animal2.Sleep() // 调用的是dog的sleep
  */
  cat := Cat{"red"}
  dog := Dog{"yellow"}
  showAnimal(&cat)
  showAnimal(&cat)

}
