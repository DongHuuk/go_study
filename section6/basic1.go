package main

import "fmt"

func main() {
	/*
		길이(가변) -> 동적으로 동작, 레퍼런스 값 타입
		슬라이스(길이 & 용량) 크기 동적으로 할당
		배열 - 길이고정, 슬라이스 - 길이 가변
		배열 - 값 타입, 슬라이스 - 참조 타입
		배열 - 복사 전달, 슬라이스 - 참조 값 전달
		배열 - 비교연산자 사용 가능, 슬라이스 - 비교 연산자 사용 불가
	*/

	//1. 배열처럼 선언, 2. make 함수: make(자료형, 길이, 용량(생략시 길이가 해당))

	var slice1 []int
	slice2 := []int{}
	slice3 := []int{1, 2, 3, 4, 5}
	slice4 := [][]int{
		{1, 2, 3, 4, 5},
		{1, 2, 3, 4, 5},
	}

	slice2 = []int{1, 2, 3}

	fmt.Printf("%-5T %d %d %v\n", slice1, len(slice1), cap(slice1), slice1)
	fmt.Printf("%-5T %d %d %v\n", slice2, len(slice2), cap(slice2), slice2)
	fmt.Printf("%-5T %d %d %v\n", slice3, len(slice3), cap(slice3), slice3)
	fmt.Printf("%-5T %d %d %v\n", slice4, len(slice4), cap(slice4), slice4)

	var slice5 []int = make([]int, 5, 10)
	var slice6 = make([]int, 6)
	slice7 := make([]int, 5, 100)
	slice8 := make([]int, 5)

	slice6[2] = 15

	fmt.Printf("%-5T %d %d %v\n", slice5, len(slice5), cap(slice5), slice5)
	fmt.Printf("%-5T %d %d %v\n", slice6, len(slice6), cap(slice6), slice6)
	fmt.Printf("%-5T %d %d %v\n", slice7, len(slice7), cap(slice7), slice7)
	fmt.Printf("%-5T %d %d %v\n", slice8, len(slice8), cap(slice8), slice8)

	var slice9 []int //nil 슬라이스(길이와 용량이 0)

	if slice9 == nil {
		fmt.Println("This is nil")
	}

}
