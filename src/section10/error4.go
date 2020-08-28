package main

import (
	"errors"
	"fmt"
	"log"
)

func notZero(i int) (string, error) {
	if i != 0 {
		s := fmt.Sprint("Hello Golang ", i)
		return s, nil
	} else {
		return "", errors.New("Not number")
	}

}

func main() {
	//Errorf를 이용한 에러 처리 - 사용자 정의 처리

	a, err1 := notZero(1)

	if err1 != nil {
		log.Fatal(err1)
	}

	fmt.Println("ex1: ", a)

	b, err2 := notZero(0)

	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println("ex1: ", b)

}
