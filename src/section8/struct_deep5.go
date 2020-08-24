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

func (e Executives) Calculator() float64{
  return e.salary + e.bonus + e.specialBonus
}

type Executives struct{
  Employee // is a 관계 = Executives is Employee
  specialBonus float64
}

func main(){
  //구조체 임베디드 메소드 오버라이딩 패턴

  e1 := Employee{"kim", 200000, 150000}
  e2 := Employee{"park", 2100000, 200000}

  ex := Executives{
    Employee{"lee", 2000000, 1000000},
    1000000,
  }
  fmt.Println("ex1 e1 : ", int(e1.Calculator()))
  fmt.Println("ex1 e2 : ", int(e2.Calculator()))

  fmt.Println("ex2 ex: ", int(ex.Calculator()))
  fmt.Println("ex2 : ", int(ex.Employee.Calculator() + ex.specialBonus ))



}
