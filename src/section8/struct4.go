package main

import "fmt"

func main(){
  //익명 구조체
  //type StructName StructType{...} <- 기본 문법

  //익명 구조체
  car1 := struct{name, color string}{"520d", "red"}

  fmt.Println(car1)
  fmt.Printf("%#v", car1)
  fmt.Println()

  cars := []struct {name, color string}{
    {"520d", "red"},
    {"530z", "black"},
    {"580a", "white"},
  }

  for _, c := range cars{
    fmt.Printf("name - %s, color - %s, struct - %#v", c.name, c.color, c)
    fmt.Println()
  }


}
