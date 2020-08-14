package main

import "fmt"

func fact(n int) int{
  if(n == 1){
    return 1
  }else{
    return n * fact(n-1)
  }
}

func prtHello(n int) {
  if(n == 0){
    return
  }else{
    fmt.Println("ex2 : (", n, ")", "hi!")
    prtHello(n - 1)
  }
}

func main(){
  //함수 고급
  //재귀 함수
  //프로그램이 보기 쉽고, 코드 간결, 오류 수정이 용어 - 장점
  //코드 이해가 어렵고, 기억공간을 많이 잡아먹는다. 무한루프의 가능성이 있다.

  x := fact(5)

  fmt.Println(x)

  prtHello(4)

}
