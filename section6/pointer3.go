package main

import "fmt"

func rptc(n *int) {
	*n = 88
}

func vptc(n int) {
	n = 88
}

func main() {
	//함수, 메서드 호출 시에 매개변수 값을 복사 전달 -> 대량 값을 복사할 경우 속도 저하 문제 발생 + 갓 변경 불가능
	//원본 값 변경 위해서 포인터로 전달
	//크기가 큰 배열의 경우 값 복사로 인해 속도 저하가 크기 때문에 포인터로 사용. (슬라이스, 맵은 참조 값 전달하므로 괜찮음)
	a := 8
	b := 8
	fmt.Println("a value - ", a)
	fmt.Println("b value - ", b)

	vptc(a)
	fmt.Println("a value - ", a)

	rptc(&b)
	fmt.Println("b value - ", b)

}
