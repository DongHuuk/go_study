package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	//대기 그룹(Wait Group)
	//실행 흐름 변경 - 고루틴이 종료 될 때까지 대기 기능
	//WaitGroup - Add(고루틴 추가), Done(작업 종료 알림), Wait(고루틴 종료시까지 대기)
	//Add로 추가 된 고루틴 개수와 Done으로 종료되는 알림 횟수가 같아야 정확하게 동작 (Add == Done)

	runtime.GOMAXPROCS(runtime.NumCPU())
	wg := new(sync.WaitGroup)

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(n int) {
			fmt.Println("WatiGroup : ", n)
			wg.Done()
		}(i)
	}

	//Add == Done이여야 한다

	wg.Wait()
	fmt.Println("WaitGroup End")
}
