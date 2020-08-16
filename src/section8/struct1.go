package main

import "fmt"

type Account struct {
  number string
  balance float64
  interest float64
}

func (a Account) Calculate() float64{
  return a.balance + (a.balance * a.interest)
}

func main(){
  //구조체 -> 구조체를 선언하고 구조체 인스턴스를 선언(리시버 방식)할 수 있다.

  kim := Account{
    number: "245-42",
    balance: 20000000,
    interest: 1.5,
  }

  choi := Account{number: "245-41", balance: 10000000, interest: 0.15}

  lee := Account{
    number:"245-42",
    balance:30000000,
  }

  park := Account{"245-43", 16000000, 0.03}

  fmt.Println("ex1 : ", kim)
  fmt.Println("ex1 : ", choi)
  fmt.Println("ex1 : ", lee)
  fmt.Println("ex1 : ", park, "")

  fmt.Println("ex2 : ", int(kim.Calculate()))
  fmt.Println("ex2 : ", int(choi.Calculate()))
  fmt.Println("ex2 : ", int(lee.Calculate()))
  fmt.Println("ex2 : ", int(park.Calculate()) , "\n")


}
