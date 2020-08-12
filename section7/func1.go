package main

import (
	"fmt"
	"strconv"
)

func helloGoLang() {
	fmt.Println("Hello Go Language")
}

func say_one(m *string) {
	fmt.Println("Hello, ", *m)
	*m = "Hello Choi Dong"
}

func sum(x int, y int) int {
	return x + y
}

func main() {
	//function
	//func 함수명(매개변수) 반환타입 or 반환 값 변수명 - 매개변수 o, 반환 값 o
	//func 함수명() 반환타입 or 반환 값 변수명 - 매개변수 x, 반환 값 o
	//func 함수명(매개변수) - 매개변수 o, 반환 값 x
	//func 함수명() - 매개변수 x, 반환 값 x
	//반환 값 다수 가능(return value)

	helloGoLang()

	str1 := "Choi"

	say_one(&str1)

	fmt.Println(str1)

	fmt.Println("sum : ", sum(5, 7))

	fmt.Printf("convertor : %T", strconv.Itoa(sum(1, 3)))

}
