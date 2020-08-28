package main

import (
	"fmt"
)

func main() {
	//Panic & Recover
	/*
		compile 언어들은 실행시 auto comiling이 동작해서 Error를 한번 잡아준다.(Java)
		Go와 같은 동적언어는 실행전 검증해준다.
		Panic - 호출시 호출한 메서드를 즉시 중지 -> defer 함수 호출 후 자기자신을 호출한 곳으로 리턴
		Runitme 이외에 사용자가 코드 흐름에 따라 에러를 발생 시킬 때 중요
		논리적 에러에 대해 처리가 가능하다
	*/

	fmt.Println("Start Main")
	panic("Error occurred : User sttoped") //방법 1
	// log.Panic("")                          //방법2
	fmt.Println("End Main") //실행 안됨

}
