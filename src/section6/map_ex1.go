package main

import "fmt"

func main() {
	/*
		레퍼런스 타입(참조 값 전달)
		key - 참조타입, value - 모든 타입
		make 함수 및 리터럴로 초기화 가능
		순서 없음 (슬라이스는 순서 있음)
	*/

	var map1 map[string]int = make(map[string]int) //정석
	var map2 = make(map[string]int)
	map3 := make(map[string]int)

	fmt.Println("ex1 : ", map1)
	fmt.Println("ex1 : ", map2)
	fmt.Println("ex1 : ", map3)
	fmt.Println()

	map4 := map[string]int{} // Json a = {} type 형태랑 비슷함
	map4["apple"] = 25
	map4["melon"] = 3829

	map5 := map[string]int{
		"apple": 15,
		"melon": 3829,
	}
	fmt.Println("ex2 : ", map4)
	fmt.Println("ex2 : ", map5)

	map6 := make(map[string]int, 10)
	map6["apple"] = 25
	map6["melon"] = 3512

	//짧은 선언은 리터럴 선언

	fmt.Println("ex2 : ", map6)
	fmt.Println("ex2 : ", map6["melon"])
}
