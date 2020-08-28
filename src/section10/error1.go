package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	//에러 처리
	/*
		에러처리 -> S/W의 품질을 향상 시키는데 가장 중요(적절한 에러 처리가 존재하지 않으면 에러 슈팅에서 소모되는 자원이 커진다.)
		ㄴ 유형코드 및 에러 정보 등을 정보를 남기는 것!
		Go에서는 error 패키지에서 error 인터페이스를 제공(하지만 Go에서는 error가 별도로 존재하지 않는다)
		type error interface { Error() string } -> 구현하여 개발자가 에러 메시지 제공
		기본적으로 메서드마다 리턴 타입이 2개(val, bool)
		1. Errorf(에러 타입 리턴) 메서드(주로)
		2. Fatal(프로그램을 종료시킴) 메서드를 통해서 에러를 출력
	*/

	//if error == nil -> 에러가 발생하지 않았다
	f, err := os.Open("error Note file") //error 발생

	if err != nil {
		log.Fatal(err.Error())
		// log.Fatal(err)
	}

	fmt.Println("ex1 : ", f.Name())

}
