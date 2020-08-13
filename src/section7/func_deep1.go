package main

import "fmt"

func multiply(n ...int) int {
	tot := 1
	for _, value := range n {
		tot *= value
	}

	return tot
}

func sum(n ...int) int {
	tot := 0
	for _, value := range n {
		tot += value
	}

	return tot
}

func prtWord(msg ...string) {
	for _, value := range msg {
		fmt.Println("ex2-prtWord : ", value)
	}
}

func main() {
	//가변 인자 실습(매개 변수 개수가 동적으로 변할 때 - 미지정)

	x := multiply(1, 2, 3, 4, 5)

	fmt.Println(x)

	y := sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(y)

	prtWord("Choi", "Dong", "Hyuk")

	a := []int{1, 2, 3, 4, 5}

	//슬라이스 대입할때 방법 알아서 넣어줌
	m := multiply(a...)
	n := sum(a...)

	fmt.Println(m)
	fmt.Println(n)

}
