package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func exe(name int) {
	r := rand.Intn(100) //랜덤 int 값 가져오기
	fmt.Println(name, " start : ", time.Now())
	for i := 0; i < 100; i++ {
		fmt.Println(name, " >>>>>>>> ", r, i)
	}
	fmt.Println(name, " end : ", time.Now())
}

func main() {
	//goroutine
	//멀티 코어(다중 cpu) 최대한 활용
	runtime.GOMAXPROCS(runtime.NumCPU())                        // runtime.NumCPU() -> 현재 내 PC Core 갯수 가져오기, GOMAXPROCS - 모든 CPU를 활용
	fmt.Println("Current System CPU : ", runtime.GOMAXPROCS(0)) // GOMAXPROCS(0) -> 현재 설정중인 GOMAXPROCS 값을 가져온다.

	fmt.Println("Main Routine Start : ", time.Now())
	for i := 0; i < 10000; i++ {
		go exe(i)
	}

	time.Sleep(8 * time.Second)
	fmt.Println("Main Routine End : ", time.Now())
}
