package main

import (
	"fmt"
)

func main() {
	//Close - 채널 닫기, 주의 -> 닫힌 채널에 값 전송 시 패닉(예외) 발생
	//Range - 채널안에서 값을 꺼낸다. (순회), 채널이 닫혀야(Close) 반복문이 종료 -> 채널이 열려 있고 값 전송이 없으면 무기한 대기

	ch := make(chan bool)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- true
		}
		close(ch) //5번 채널에 값 전송 후 채널 닫기
	}()

	fmt.Printf("%T, %v", ch, ch)

	for i := range ch {
		fmt.Println("ex1: ", i)
	}

}
