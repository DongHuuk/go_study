package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a interface{} = 15
	b := a.(int)
	c := b

	fmt.Println("ex1: ", a)
	fmt.Println("ex1: ", reflect.TypeOf(a))
	fmt.Println("ex1: ", b)
	fmt.Println("ex1: ", reflect.TypeOf(b))
	fmt.Println("ex1: ", c)
	fmt.Println("ex1: ", reflect.TypeOf(c))
	fmt.Printf("ex1: %T", c)
	fmt.Println()
	//TypeOf == %T

	//저장된 실제 타입 검사
	if v, ok := a.(int); ok {
		fmt.Println("ex2: ", v, ok)
	}
}
