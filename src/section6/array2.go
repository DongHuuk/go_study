package main

import "fmt"

func main() {
	arr1 := [5]int{1, 10, 100, 1000, 10000}

	for i := 0; i < len(arr1); i++ {
		fmt.Println("ex: 1 ", arr1[i])
	}

	arr2 := [5]int{7, 77, 777, 7777, 77777}

	//range 중요
	for i, v := range arr2 {
		fmt.Println("ex: 2 : ", i, v)
	}

	fmt.Println()

	for _, v := range arr2 {
		fmt.Println("ex: 2 : ", v)
	}

	fmt.Println()

	for v := range arr2 {
		fmt.Println("ex: 3 ", arr2[v])
	}

}
