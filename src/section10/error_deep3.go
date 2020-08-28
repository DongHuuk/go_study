package main

import (
	"fmt"
	"log"
	"math"
	"time"
)

//예외 처리 구조체
type PowError struct {
	time    time.Time //에러 발생 시간
	value   float64   //interface{} 를 이용해서 모든 타입 대응 가능
	message string
}

func (p PowError) Error() string {
	return fmt.Sprintf("[%v]Error - Input Value(value : %g) - %s", p.time, p.value, p.message) // v = time
}

func Power(f, i float64) (float64, error) {
	if f == 0 {
		return 0, PowError{time.Now(), f, "0은 처리할 수 없습니다."}
	}

	if i == 0 {
		return 1, PowError{time.Now(), i, "0은 처리할 수 없습니다."}
	}

	if math.IsNaN(f) {
		return 0, PowError{time.Now(), f, "숫자가 아닙니다."}
	}

	if math.IsNaN(i) {
		return 0, PowError{time.Now(), i, "숫자가 아닙니다."}
	}

	return math.Pow(f, i), nil
}

func main() {
	//error 타입이 아닌 경우 에러 처리 방법
	//Error 메서드를 구현해서 사용자 정의 Error 구조체 생성 후 처리

	v, err1 := Power(10, 3)

	if err1 != nil {
		log.Fatal(err1.Error())
	}

	fmt.Println("ex1 : ", v)

	t, err2 := Power(0, 3)

	if err2 != nil {
		// log.Fatal(err2.Error())
		fmt.Println(err2.(PowError).Error())
		fmt.Println(err2.(PowError).time)
		fmt.Println(err2.(PowError).value)
		fmt.Println(err2.(PowError).message)
	}
	
	fmt.Println("ex1 : ", int(t))
}
