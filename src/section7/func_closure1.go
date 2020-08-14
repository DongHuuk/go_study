package main

import "fmt"

func main(){
  /*
    익명함수 사용할 경우 함수를 변수에 할당해서 사용 가능
    -> 함수 내에서 함수를 정의 및 선언 가능(외부 함수에 선언 된 변수에 접근 가능한 함수)
    함수가 선언되는 순간에 함수가 실행 될 때 실체의 외부 변수에 접근하기 위한 스냅샷(객체)
  */

  multiply := func(x, y int) int{
    return x * y
  }

  r1 := multiply(5, 10)
  fmt.Println("ex: 1 ", r1)


  m, n := 5, 10

  sum := func(c int) int {
    return m + n + c // 지역 변수 소멸 x (호출 시 마다 사용 가능)
  }

  r2 := sum(10)

  fmt.Println("ex2 : ", r2)


}
