package main

import "fmt"

type Car struct{//대문자 - public
  name string "차량명"
  color string "색상"
  company string "제조사"
  detail spec "상세"
}

type spec struct{//소문자 - private
  length int "전장"
  height int "전고"
  width int "전축"
}


func main(){
  car1 := Car{
    "520d",
    "sliver",
    "bmw",
    spec{100, 200, 300},
  }

  fmt.Println("ex1: ", car1.name)
  fmt.Println("ex1: ", car1.color)
  fmt.Println("ex1: ", car1.company)
  fmt.Printf("ex1: %#v", car1.detail)
  fmt.Println()
  fmt.Println("ex1: ", car1.detail.length)
  fmt.Println("ex1: ", car1.detail.height)
  fmt.Println("ex1: ", car1.detail.width)

}
