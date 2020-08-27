package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	/*
		뮤텍스 - 상호 배제 -> Thread(GoRouten)들이 서로 running time에 서로 영향을 주지 않게 실행되게 하는 기술
		(단독으로 실행되게 하는 기술, 무결성을 보장)
		쓰기 전용, 읽기 전용 뮤텍스가 존재한다.
	*/

	//쓰기 읽기 동작 순서가 일정하지 않아 잘못된 오류를 반환 할 가능성 증가
	runtime.GOMAXPROCS(runtime.NumCPU())
	data := 0

	go func() {
		for i := 0; i < 10; i++ {
			data++
			fmt.Println("Write data - ", data)
			time.Sleep(200 * time.Millisecond)
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("Read1 data - ", data)
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("Read2 data - ", data)
			time.Sleep(1 * time.Second)
		}
	}()
	time.Sleep(5 * time.Second)
}
