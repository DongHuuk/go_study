package main

import (
	"fmt"
	"runtime"
)

func main() {
	//비동기 - 버퍼 사용
	runtime.GOMAXPROCS(1)
	ch := make(chan bool, 2) // 용량 2 설정
	cnt := 12

	go func() {
		for i := 0; i < cnt; i++ {
			ch <- true //버퍼 용량이 가득 차면 대기
			fmt.Println("Go : ", i)
		}
	}()

	for i := 0; i < cnt; i++ {
		<-ch //버퍼 용량의 값을 가져오고, 버퍼가 비어있으면 대기
		fmt.Println("Main : ", i)
	}
}
