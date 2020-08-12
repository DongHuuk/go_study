package main

import "fmt"

func main() {
	//슬라이스 복사
	//copy(복사 대상, 원본)
	//make로 공간 할당 후 복사해야지 가능

	s1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s2 := make([]int, 5)
	s3 := make([]int, 10)
	s4 := []int{}

	copy(s2, s1)
	copy(s3, s1)
	copy(s4, s1)

	fmt.Printf("%v, len = %d, cap = %d\n", s2, len(s2), cap(s2))
	fmt.Printf("%v, len = %d, cap = %d\n", s3, len(s3), cap(s3))
	fmt.Printf("%v, len = %d, cap = %d\n", s4, len(s4), cap(s4))

	//참조하지 않는다라는 증거
	a := []int{1, 2, 3, 4, 5}
	b := make([]int, 5)
	copy(b, a)
	b[0] = 7
	b[4] = 10
	fmt.Println()
	fmt.Println("ex2: ", a)
	fmt.Println("ex2: ", b)

	//슬라이스 추출은 참조로 값이 가져온다. ex) c[0:3]
	c := [5]int{1, 2, 3, 4, 5}
	d := c[0:3]
	d[1] = 7
	fmt.Println()
	fmt.Println("ex3: ", c)
	fmt.Println("ex3: ", d)

	A := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//세번째는 cap을 늘려준다.
	B := A[0:5:7]

	fmt.Println("ex4 : ", len(B), cap(B), B)
}
