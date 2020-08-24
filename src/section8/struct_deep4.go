//임베디드 패턴(Overide)
package main

import "fmt"

type Employee struct{
  name string
  salary float64
  bonus float64
}

func (e Employee) Calculator() float64{
  return e.salary + e.bonus
}

type Executives struct{
  Employee // is a 관계 = Executives is Employee
  specialBonus float64
}

func main(){
  //구조체 임베디드 패턴
  //다른 관점으로 메소드를 재사용하는 장점 제공
  //상속을 허용하지 않는 Go 언어에서 메소드 상속 활용을 위한 패턴

  e1 := Employee{"kim", 200000, 150000}
  e2 := Employee{"park", 2100000, 200000}

  ex := Executives{
    Employee{"lee", 20000000, 1500000},
    1000000,
  }
  fmt.Println("ex1 e1 : ", int(e1.Calculator()))
  fmt.Println("ex1 e2 : ", int(e2.Calculator()))

  fmt.Println("ex1 ex: ", int(ex.Calculator() + ex.specialBonus))

}
