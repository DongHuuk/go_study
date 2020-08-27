package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	//원자성 사용(atomic) -> 기능적으로 분할 불가능한 완전 보증된 일련의 조작(commit & rollback)
	//모든 조작이 완료 될 때까지 다른 프로세스는 개입이 불가능
	//sync/atomic에서 원자적 연산자 제공
	//https://golang.org/pkg/sync/atomic에서 확인 가능
	//주로 공용 변수에 관란 계산시 사용

	runtime.GOMAXPROCS(runtime.NumCPU())

	var cnt int64 = 0

	wg := new(sync.WaitGroup)

	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func(n int) {
			// cnt++
			atomic.AddInt64(&cnt, 1)
			wg.Done()
		}(i)
	}

	for i := 0; i < 2000; i++ {
		wg.Add(1)
		go func(n int) {
			// cnt--
			atomic.AddInt64(&cnt, -1)
			wg.Done()
		}(i)
	}

	//Add(7000) == Done(7000)
	wg.Wait()

	finalCnt := atomic.LoadInt64(&cnt)
	fmt.Println("WaitGroup End >>>>> ", cnt)
	fmt.Println("WaitGroup End >>>>> ", finalCnt)
}
