package main

import (
	"fmt"
	"time"
)

func work1(v chan int) {
	fmt.Println("Work1 : Start --- > ", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println("Work1 : End --- > ", time.Now())
	v <- 1
}

func work2(v chan int) {
	fmt.Println("Work2 : Start --- > ", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println("Work2 : End --- > ", time.Now())
	v <- 2
}

func main() {
	/*
		고루틴(GoRoutine)간의 상호 정보(데이터) 교환 및 실행 흐름 동기화를 위해 사용 - 채널은 동기식 방식이 디폴트
		실행 흐름 제어 가능(동기, 비동기) -> 일반 변수로 선언 후 사용이 가능하다
		데이터 전달 자료형 선언 후 사용(지정 된 타입만 주고받을 수 있음)
		interface{} 사용해서 타입 캐스팅해서 사용 가능
		값을 전달(복사 후 : bool, int 등), 포인터(슬라이스, 맵) 등을 전달시에는 주의! -> 동기화 사용(Mutex)
		멀티프로세싱 처리에서 교착상대(경쟁) 주의!
		<-, -> 사용 ex) (채널 <- 데이터) - 송신, (변수 -> 채널) - 수신(할당)
	*/

	fmt.Println("Main : Start ---> ", time.Now())

	// var v chan int
	// v = make(chan int)
	v := make(chan int) // int형 채널 선언

	go work1(v)
	go work2(v)

	<-v // 동기

	fmt.Println("Main : End ---> ", time.Now())
}
