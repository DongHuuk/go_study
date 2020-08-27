package main

import (
	"fmt"
	"runtime"
	"sync"
)

//구조체 - 공유 데이터
type count struct {
	number int
	mutex  sync.Mutex
}

//c로 접근할거라 구조체 내에 선언
func (c *count) increment() {
	c.mutex.Lock()
	c.number += 1
	c.mutex.Unlock()
}

func (c *count) result() {
	fmt.Println("ex1: ", c.number)
}

func main() {
	/*
		실행 흐름 제어 및 변수 동기화 가능
		공유 데이터 보호가 가장 중요 (memory safety)
		뮤텍스(mutex) - 공유 데이터의 동기화
		sync.Mutex 선언 후 Lock, UnLock 사용
	*/

	runtime.GOMAXPROCS(runtime.NumCPU())

	c := count{number: 0}

	done := make(chan bool)

	for i := 1; i <= 10000; i++ {
		go func() {
			c.increment()
			done <- true
			runtime.Gosched() //다른 CPU에게 양보
		}()
	}

	for i := 1; i <= 10000; i++ {
		<-done
	}

	c.result()
}
