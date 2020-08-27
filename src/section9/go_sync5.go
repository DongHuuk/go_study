package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	//실행 흐름을 조절
	/*
		동기화 상태(조건) 메서드 사용
		Wait(일시정지), notify, notifyAll(wake up all) - 타 언어
		Wait, Signal, Broadcas(wake up all) - Golang
	*/
	runtime.GOMAXPROCS(runtime.NumCPU())

	var mutex = new(sync.Mutex)
	var condition = sync.NewCond(mutex)

	ch := make(chan int, 5) //비동기 버퍼 size 5

	for i := 0; i < 5; i++ {
		go func(n int) {
			mutex.Lock()
			ch <- 777
			fmt.Println("Goroutine Wating : ", n)
			condition.Wait() // goroutine 일시정지
			fmt.Println("Wating End", n)
			mutex.Unlock()
		}(i)
	}

	for i := 0; i < 5; i++ {
		fmt.Println("received : ", <-ch)
	}

	// for i := 0; i < 5; i++ {
	// 	mutex.Lock()
	// 	fmt.Println("Wake Goroutine(Signal) : ", i)
	// 	condition.Signal() //모든 goroutine 생성 후 동작 한개씩 wake up
	// 	mutex.Unlock()
	// }
	condition.Signal()
	mutex.Lock()
	fmt.Println("Wake up Goroutine All(BroadCast)")
	condition.Broadcast()
	mutex.Unlock()

	time.Sleep(2 * time.Second)

}
