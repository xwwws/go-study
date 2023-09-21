package main

import "fmt"

type Book struct {
  title string
  auth string
}
//  struct  值传递

func changeBook1(book Book)  {
  book.auth = "bookauth"
}
func changeBook2(book *Book)  {
  book.auth = "bookauth"
}

func main()  {
  var book1 Book
  book1.title = "Golang"
  book1.auth = "zhang3"

  fmt.Printf("%v\n", book1)

  changeBook1(book1)
  fmt.Printf("%v\n", book1)
  changeBook2(&book1)
  fmt.Printf("%v\n", book1)


}
