package main

import (
  "fmt"
)

func sayHello(str string){
  defer func(){
    fmt.Println("defer : ", str)
  }()

  func(){
    fmt.Println("Hi! ")
  }()

  // str = "change str"  변경된 값이 Defer에 적용됨

}

func main(){
    sayHello("Golang")
}
