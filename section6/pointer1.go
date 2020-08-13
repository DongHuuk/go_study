package main

import "fmt"

func main() {

	//Go - 포인터 지원
	//변수의 ㅈ역성, 연속된 메모리 참조 ...
	//복잡 연산 같은건 사용이 불가능 (주소의 값은 직접 변경 불가능하다. 코딩의 실수 방지)
	//*(에스터리스크)로 사용
	//nul로 초기화 (nil != 0)

	var a *int            //초기화 x
	var b *int = new(int) //초기화 o

	// fmt.Println(a)
	// fmt.Println(b)

	i := 8

	a = &i
	b = &i
	c := &i

	fmt.Println("this is i's value - ", &i)
	fmt.Println("this is i's value - ", i)

	fmt.Println("print a, b, c values")
	fmt.Println(&a)
	fmt.Println(*a)
	fmt.Println(a)

	fmt.Println(&b)
	fmt.Println(*b)
	fmt.Println(b)

	fmt.Println(&c)
	fmt.Println(*c)
	fmt.Println(c)

	*b = 10

	fmt.Println("print a, b, c values")
	fmt.Println(&a)
	fmt.Println(*a)
	fmt.Println(a)

	fmt.Println(&b)
	fmt.Println(*b)
	fmt.Println(b)

	fmt.Println(&c)
	fmt.Println(*c)
	fmt.Println(c)
	fmt.Println()

	fmt.Println("i value")
	fmt.Println(i)
}
