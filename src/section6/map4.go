package main

import "fmt"

func main() {
	//Map 조회 시 주의사항
	//Map 조회 시 값이 존재하지 않으면 해당 타입의 기본 값으로 리턴

	map1 := map[string]int{
		"apple":  15,
		"banana": 25,
		"melon":  39273,
		"lemon":  0,
	}
	value1 := map1["lemon"]
	value2 := map1["kiwi"]
	value3, ok := map1["kiwi"]
	_, value4 := map1["lemon"]

	fmt.Println(value1)
	fmt.Println(value4)
	fmt.Println(value2)
	fmt.Println(value3)
	fmt.Println(ok)

	if _, ok := map1["solaly"]; ok {
		fmt.Println("have solaly")
	} else {
		fmt.Println("dont have solaly")
	}
}
