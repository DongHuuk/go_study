package main

import (
	"errors"
	"fmt"
)

func main() {
	//errors 패키지의 new 메서드를 이용한 에러 생성
	//Errorf <<< new Error

	var err1 error = errors.New("Not number - 1")
	err2 := errors.New("Not number - 2")

	fmt.Println(err1)
	fmt.Println(err1.Error())
	fmt.Println(err2)
	fmt.Println(err2.Error())
}
