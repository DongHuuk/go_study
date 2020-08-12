package main

import "fmt"

func multiply(x, y int) (int, int) {
	return x * 10, y * 10
}

func arrayMultiply(a, b, c, d, e int) (int, int, int, int, int) {
	return a * 1, b * 2, c * 3, d * 4, e * 5
}

func main() {
	//다중 값 반환
	a, b := multiply(3, 5)
	c, _ := multiply(10, 5)
	_, d := multiply(10, 5)

	fmt.Println(a, b, c, d)

	q, w, e, r, t := arrayMultiply(1, 2, 3, 4, 5)
	var p, _, o, _, i int = arrayMultiply(1, 2, 3, 4, 5)

	fmt.Println("ex2 : ", q, w, e, r, t)
	fmt.Println("ex2 : ", p, o, i)

}
