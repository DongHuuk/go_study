package main

import (
  "fmt"
)

func a(){
  //end만 지연 적용된다 (내부함수는 적용안됨)
  defer end(start("b"))
  fmt.Println("in a")
}

func start(str string) string{
  fmt.Println("in start : ", str)
  return str
}

func end(str string){
  fmt.Println("in end : ", str)
}

func main(){
  a()
}
