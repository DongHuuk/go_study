package main

import (
	"fmt"
	"time"
)

func main() {
	//채널 Select -> 채널에 값이 수신되면 해당 case문이 실행이 된다
	//일회성 구문이므로 for를 이용한다
	//default 처리 주의

	ch1 := make(chan int)
	ch2 := make(chan string)

	go func() {
		for {
			ch1 <- 77
			time.Sleep(250 * time.Millisecond)
		}
	}()

	go func() {
		for {
			ch2 <- "Golang"
			time.Sleep(500 * time.Millisecond)
		}
	}()

	go func() {
		for {
			select {
			case num := <-ch1:
				fmt.Println("get number by ch1 - ", num)
			case str := <-ch2:
				fmt.Println("get string by ch2 - ", str)
			} //채널을 수,발신하는 용도로 사용하기 위해선 default를 사용해서는 안된다.
		}
	}()

	time.Sleep(7 * time.Second)
}
