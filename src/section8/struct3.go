package main

import "fmt"

type Car struct{
  color string
  name string
}

func main(){
  c1 := Car{"red", "220d"}
  c2 := new(Car)
  c2.color, c2.name = "white", "sonata"
  c3 := &Car{}
  c4 := &Car{"black", "520d"}

  fmt.Println(c1, c2, c3, c4)

}
