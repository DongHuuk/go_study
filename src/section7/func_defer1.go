package main

import "fmt"

func ex_f1() {
  fmt.Println("f1 : start")
  defer ex_f2()
  fmt.Println("f1 : end")
}

func ex_f2(){
  fmt.Println("f2 : called")
}


func main(){
  //finally와 유사
  //Defer를 호출한 함수가 종료하기 직전에 호출
  //주로 리소스를 반환, 열린 파일 닫기, Mutex 잠금 해제

  ex_f1()

}
