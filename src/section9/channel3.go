package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan bool)
	cnt := 6

	go func() {
		for i := 0; i < cnt; i++ {
			ch <- true
			fmt.Println("Go : ", i)
			time.Sleep(100 * time.Millisecond)
		}
	}()
	//chan은 동기식이니깐 <- 송신을 하면 수신을 할 때까지 대기했다가 수신이 되면 다시 동작하는 방식의 고루틴
	for i := 0; i < cnt; i++ {
		<-ch
		fmt.Println("Main : ", i)
	}
}
