package main

import (
	"fmt"
	"time"
)

func exe1() {
	fmt.Println("exe1 func start", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println("exe1 func end", time.Now())
}

func exe2() {
	fmt.Println("exe2 func start", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println("exe2 func end", time.Now())
}

func exe3() {
	fmt.Println("exe3 func start", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println("exe3 func end", time.Now())
}

func main() {
	/*
		고루틴(goroutine)
		쓰레기(Thread)와 유사하게 동작
		생성 방법이 간단하며 리소스를 적게 차지
		수많은 고루팅 동시 생성 실행 가능
		비동기적 함수 루틴 실행() -> 채널을 통한 통신 가능
		공유메모리 사용 시에 정확한 동기화 코딩 필요
		싱글루틴(main thread)에 비해 항상 빠른 처리 결과는 아니다. -> 고루틴이 한개(main만)있는 것보다 반드시 빠른 속도를 보장하지 않는다.
	*/
	/*
		멀티 쓰레드 시스템 개발시 장점
		 - 응답성 향상, 자원 공유를 효율적으로 활용 가능, 작업이 분리되어 코드가 간결화
		 멀티 쓰레드 시스템 개발시 단점
		 - 구현하기 어렵다. 테스트 및 디버깅 어려움, 전체 프로세스에 악영향, 성능 저하, 동기화 코딩 반드시 숙지, 데드락(교착상태)...
	*/

	exe1() // 먼저 실행
	fmt.Println("Main Routine Start", time.Now())
	go exe2()
	go exe3()
	time.Sleep(5 * time.Second)
	fmt.Println("Main Routine End", time.Now())
}
