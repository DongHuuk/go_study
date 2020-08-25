package main

import (
	"fmt"
)

type Account struct {
	name string
	age  int32
}

func checkType(arg interface{}) {
	switch arg.(type) {
	case bool:
		fmt.Println("This is boolean ", arg)
	case int, int8, int16, int32, int64:
		fmt.Println("This is Integer ", arg)
	case float32, float64:
		fmt.Println("This is float ", arg)
	case string:
		fmt.Println("This is string ", arg)
	case nil:
		fmt.Println("This is nil ", arg)
	default:
		fmt.Println("Type Error ", arg)
	}
}

func chekcType_Object(arg interface{}) {
	switch arg.(type) {
	case Account:
		fmt.Println("This is Account")
	case *Account:
		fmt.Println("This is pointer Account")
	default:
		fmt.Println("Type Error")
	}
}

func main() {
	//실제 타입 검사 switch <- 자주 사용함
	//인터페이스로 자료를 받아서 타입체크 후 형 변환 한 후 사용이 가능하다

	checkType(true)
	checkType(1)
	checkType(22.5)
	checkType("Hello")
	checkType(nil)
	a := Account{"name", 25}
	chekcType_Object(a)
	b := &Account{"name", 25}
	chekcType_Object(b)

}
