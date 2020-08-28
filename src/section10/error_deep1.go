package main

import (
	"errors"
	"fmt"
	"math"
)

func number(f, i float64) (float64, error) {
	if f == 0 {
		return 0, errors.New("0은 사용불가")
	}

	if i == 0 {
		return 1, errors.New("0제곱은 사용불가")
	}

	return math.Pow(f, i), nil
}

func main() {
	//Go errors패키지 New 메서드 사용 -> 사용자 에러 처리 생성

	if f, err := number(8, 3); err != nil {
		fmt.Printf("Error Message : %s\n", err)
	} else {
		fmt.Println("ex1 : ", f)
	}

	if f, err := number(0, 3); err != nil {
		fmt.Printf("Error Message : %s\n", err)
	} else {
		fmt.Println("ex1 : ", f)
	}

	if f, err := number(5, 0); err != nil {
		fmt.Printf("Error Message : %s\n", err)
	} else {
		fmt.Println("ex1 : ", f)
	}
}
