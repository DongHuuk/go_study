package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	/*
		뮤텍스 - 상호 배제 -> Thread(GoRouten)들이 서로 running time에 서로 영향을 주지 않게 실행되게 하는 기술
		(단독으로 실행되게 하는 기술, 무결성을 보장)
		쓰기 전용, 읽기 전용 뮤텍스가 존재한다.
		쓰기 뮤텍스 - RWMutex -> 쓰기 시도 중에는 다른 곳에서 값을 읽으면 안됨, 읽기 및 쓰기 전부 락
		읽기 뮤텍스 - RMutex -> 읽기 시도 중 값 변경 방지, 쓰는 작업만 락
	*/

	//쓰기 읽기 동작 순서가 일정하지 않아 잘못된 오류를 반환 할 가능성 증가
	//쓰기 전용 및 읽기 전용 사용
	runtime.GOMAXPROCS(runtime.NumCPU())

	data := 0
	mutex := new(sync.RWMutex) // var mutex = new(sync.RWMutex)

	go func() {
		for i := 0; i < 10; i++ {
			mutex.Lock()
			data++
			fmt.Println("Write data - ", data)
			time.Sleep(200 * time.Millisecond)
			mutex.Unlock()
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			mutex.RLock()
			fmt.Println("Read1 data - ", data)
			time.Sleep(1 * time.Second)
			mutex.RUnlock()
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			mutex.RLock()
			fmt.Println("Read2 data - ", data)
			time.Sleep(1 * time.Second)
			mutex.RUnlock()
		}
	}()

	time.Sleep(12 * time.Second)
}
