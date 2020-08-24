package main

import "fmt"

type Account struct{
  number string
  balance float64
  interest float64
}

//리시버 형식으로 선언
func (a Account) CalculatorD(bonus float64){
  a.balance = a.balance + (a.balance * a.interest) + bonus
}

func (a *Account) CalculatorP(bonus float64){
  a.balance = a.balance + (a.balance * a.interest) + bonus
}

func main(){
  //구조체 인스턴스 값 변경 시 포인터로 전달해야 한다.
  //대게의 경우 값을 전달한다.

  a := Account{"245-900", 10000000, 0.15}
  b := Account{"245-901", 17500000, 0.1}

  fmt.Println("ex1 a : ", a)
  fmt.Println("ex1 b : ", b)

  //리시버로 선언했을 경우 &를 붙이지 않고 그냥 넘겨도 값의 변경이 적용이 된다(포인터 사용)
  a.CalculatorD(1000000)
  b.CalculatorP(1000000)

  fmt.Println("ex1 a : ", int(a.balance))
  fmt.Println("ex1 b : ", int(b.balance))

}
