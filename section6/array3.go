package main

import "fmt"

func main() {
	arr1 := [5]int{1, 2, 3, 4, 5}
	arr2 := arr1

	fmt.Println(arr1, &arr1)
	fmt.Println(arr2, &arr2)
	fmt.Println(arr1 == arr2)

	//v = original value, p = pointer
	fmt.Printf("ex1 : %p %v\n", &arr1, arr1)
	fmt.Printf("ex2 : %p %v\n", &arr2, arr2)

	fmt.Println()
	arr1[2] = 100

	fmt.Printf("ex1 : %p %v\n", &arr1, arr1)
	fmt.Printf("ex2 : %p %v\n", &arr2, arr2)
}
