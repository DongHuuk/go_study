package main

import "fmt"

//alias
type cnt int

func main(){

  a := cnt(15)
  fmt.Println("ex1 : ", a)

  var b cnt = 15
  fmt.Println("ex1 : ", b)


  testConverT(int(a)) // 형변환
  testConverD(a)
}

func testConverT(i int){
  fmt.Println("ex3 : (Default Type) : ", i)
}

func testConverD(i cnt){
  fmt.Println("ex4 : (Custom Type) : ", i)
}
