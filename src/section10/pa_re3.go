package main

import (
	"fmt"
)

func simpleRun() {
	defer func() {
		if s := recover(); s != nil {
			fmt.Println("function inner defer - ", s)
		}
	}()

	a := [3]int{1, 2, 3} //array
	// b := []int{}        slice

	for i := 0; i < 5; i++ {
		fmt.Println("ex1 : ", a[i])
	}

}

func simpleStr() string {
	return fmt.Sprint("func inner Error")
}

func main() {
	//Recover - 에러 복구 가능
	/*
		Panic으로 발생한 에러를 복구 후 코드를 재 실행(종료 x)
		코드 흐름을 정상 상태로 복구하는 기능
		Panic에서 설정한 메시지를 받아올 수 있다.
	*/

	simpleRun()

	fmt.Println("Mian func")

}
