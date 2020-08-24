package main

import "fmt"

type Account struct{
  number string
  balance float64
  interest float64
}

func CalculatorD(a Account){
  a.balance = a.balance + (a.balance * a.interest)
}

func CalculatorP(a *Account){
  a.balance = a.balance + (a.balance * a.interest)
}

func main(){

  a := Account{"245-900", 10000000, 0.15}
  b := Account{"245-901", 17500000, 0.1}

  fmt.Println("ex1 a : ", a)
  fmt.Println("ex1 b : ", b)

  CalculatorD(a)
  CalculatorP(&b)

  fmt.Println("ex1 a : ", int(a.balance))
  fmt.Println("ex1 b : ", int(b.balance))

}
