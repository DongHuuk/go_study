package main

import "fmt"

func multiply(x, y int) (r int) {
	r = x * y
	return r
}

func sum(x, y int) (r int) {
	r = x + y
	return r
}

func main() {
	//함수를 변수에 할당

	f := []func(int, int) int{multiply, sum}

	a := f[0](2, 3)
	b := f[0](3, 4)

	fmt.Println(a)
	fmt.Println(b)

	var f1 func(int, int) int = multiply
	f2 := sum

	fmt.Println(f1(1, 20))
	fmt.Println(f2(1, 20))

	m := map[string]func(int, int) int{
		"mul_func": multiply,
		"sum_func": sum,
	}

	fmt.Println("ex3 : ", m["mul_func"](1, 10))
	fmt.Println("ex3 : ", m["sum_func"](1, 10))

}
