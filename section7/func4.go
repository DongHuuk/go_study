package main

import "fmt"

func multiply(x, y int) (r1 int, r2 int) {
	r1 = x * 10
	r2 = y * 20
	return r1, r2
}

func main() {
	fmt.Println(multiply(5, 6))
}
