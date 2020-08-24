package main

import "fmt"

type Account struct{
  number string
  balance float64
  interest float64
}

//생성자패턴 - 패턴형식으로 많이 사용하고 있음
func NewAccount(number string, balance, interest float64) *Account{ //포인터 반환이 아닌 경우 값 복사
  return &Account{number, balance, interest} //구조체 인스턴스를 생성 한 뒤 포인터를 리턴
}

func main(){
  //Account -> 생성자
  a := Account{"201-900", 10000000, 0.05}
  var b *Account = new(Account)
  b.number = "201-901" //getter, setter
  b.balance = 20000000
  b.interest = 0.15

  fmt.Println("ex1 : ", a)
  fmt.Println("ex1 : ", *b)

  //이러한 생성자 패턴을 사용하면 Java에서의 생성자처럼 선언이 가능하다.
  c := NewAccount("201-902", 15000000, 0.04)

  fmt.Println(*c)

}
