package main

import (
	"fmt"
)

func simpleRun() {
	defer func() {
		s := recover()
		fmt.Println("function inner defer - ", s)
	}()
	panic(simpleStr()) //이 아래의 동작들은 recover 동작 후에도 동작하지 않는다. -> Panic 동작 후 이 함수를 호출한 곳으로 즉시 이동하기 때문
	fmt.Println("function simpleRun")
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
