package main

import "fmt"

func main(){
  //익명 함수(Anonymous Function)
  //즉시 실행(선언과 동시)

  func() {
    fmt.Println("ex1 : Anonymous Function")
  }()

  msg := "Hello goLang"

  func(str string){
    fmt.Println("ex2 : ", str)
  }(msg)

  func(x, y int){
    fmt.Println("ex3 : ", x + y)
  }(3, 5)

  f := func(x, y int) int {
    return x * y
  }(10, 20)

  fmt.Println("ex4 : ", f)

}
