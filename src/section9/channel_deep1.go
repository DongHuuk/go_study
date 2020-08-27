package main

import (
	"fmt"
	"time"
)

func sendOnly(c chan<- int, cnt int) {
	for i := 0; i < cnt; i++ {
		c <- i
	}
	c <- 777
}
func receiveOnly(c <-chan int) {
	for i := range c {
		fmt.Println("received : ", i)
	}

	fmt.Println(<-c)
}

func main() {
	/*
		함수 등의 매개변수에 수신 및 발신 전용 채널을 지정 가능
		전용 채널 설정 후 방향이 다를 경우 예외 발생
		발신 전용 - chan <- type
		수신 전용 - <- chan type
		매개 변수를 통해 전용 채널을 확인 가능
		채널 또한 함수의 반환 값으로 사용이 가능
	*/

	c := make(chan int)

	go sendOnly(c, 10) //발신
	go receiveOnly(c)  //수신

	time.Sleep(2 * time.Second)
}
