package main

import "fmt"

type Car struct{
  name string
  color string
  price int64
  tax int64
}

//매개변수가 구조체 (일반 메서드)
func Price(c Car) int64{
  return c.price + c.tax
}

//구조체 <-> 메소드 바인딩
func (c Car) Price() int64{
  return c.price + c.tax
}

func main(){
  /*
  객체 지향 타입을 구조체로 정의한다. (클래스, 상속 개념 없음)
  객체 지형은 클래스(속성, 기능(상태 - 메소드))로 표현 -> 재사용성과 관리성, 신뢰성이 높은 프로그래밍
  GO는 객체지향의 특징을 가지고 있지는 않지만, 인터페이스를 이용해서 다형성을 지원하고 구조체로 클래스를 구현이 가능
  상태, 메소드 - 분리해서 정의(결합성 없음)
  사용자 정의 타입 - 구조체, 인터페이스, 기본 타입(int float string...), 함수 재정의
  */

  var carTest Car = Car{}
  carTest.name = "520d"

  bent := Car{
    name : "220d",
    price : 5000000,
    color : "black",
    tax : 129321,
  }

  bmw := Car{
    name : "520d",
    price : 7000000,
    color : "black",
    tax : 123129321,
  }

  fmt.Println("ex1 : ", bent, &bent)
  fmt.Println("ex1 : ", bmw, &bmw)

  fmt.Println("ex2 : ", Price(bmw))
  fmt.Println("ex2 : ", Price(bent))

  fmt.Println("ex3 : ", bmw.Price())
  fmt.Println("ex3 : ", bent.Price())
}
