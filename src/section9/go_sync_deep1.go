package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

//한번만 실행될 코드
func onceTest() {
	fmt.Println("init Code")
}

func main() {
	//Once - 특정 고루틴에서 단 한번만 실행(주로 초기화에 사용), Do로 실행
	runtime.GOMAXPROCS(runtime.NumCPU())

	once := new(sync.Once)

	for i := 0; i < 5; i++ {
		go func(n int) {
			fmt.Println("Goroutine : ", n)
			once.Do(onceTest)
		}(i)
	}

	time.Sleep(2 * time.Second)
}
