package main

import "fmt"

func sum(i int, f func(int, int)) {
	f(i, 10)
}

func add(x, y int) {
	fmt.Println("ex1 : ", x+y)
}

func multi_value(i int) {
	i = i * 10
}

func multi_reference(i *int) {
	*i *= 10
}

func main() {
	//함수 전달(콜백), 참조 전달(call by reference), 값 전달(call by value)

	sum(100, add)

	//call by value
	a := 100
	multi_value(a)
	fmt.Println("ex 2 : ", a)

	//call by reference
	b := 100
	multi_reference(&b)
	fmt.Println("ex 2: ", b)

}
