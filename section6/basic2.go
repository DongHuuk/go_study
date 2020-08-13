package main

import "fmt"

func main() {
	//슬라이스 참조타입 증명
	var slice1 []int = []int{1, 2, 3, 4, 5}
	slice2 := slice1

	var arr1 [5]int = [5]int{1, 2, 3, 4, 5}
	arr2 := arr1

	fmt.Printf("%-5T %p %v\n", slice1, &slice1, slice1)
	fmt.Printf("%-5T %p %v\n", slice2, &slice2, slice2)
	fmt.Printf("%-5T %p %v\n", arr1, &arr1, arr1)
	fmt.Printf("%-5T %p %v\n", arr2, &arr2, arr2)

	fmt.Println()

	arr1[1] = 9999
	slice1[1] = 9999

	fmt.Printf("%-5T %p %v\n", slice1, &slice1, slice1)
	fmt.Printf("%-5T %p %v\n", slice2, &slice2, slice2)
	fmt.Printf("%-5T %p %v\n", arr1, &arr1, arr1)
	fmt.Printf("%-5T %p %v\n", arr2, &arr2, arr2)

	slice3 := make([]int, 50, 100)
	fmt.Println("ex 3 ", slice3[5])
	// fmt.Println("ex 3 ", slice3[50]) index over error
	fmt.Println()

	for i, v := range slice2 {
		fmt.Println("ex 3 : ", i, v)
	}

}
