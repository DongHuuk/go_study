package main

import "fmt"

func fnV(arr [4]int) {
	for i := range arr {
		arr[i] = 0
	}
}

func fnP(arr *[4]int) {
	for i := range arr {
		arr[i] = 0
	}
}

func main() {
	fmt.Println("not run")
	arr1 := [4]int{1, 2, 3, 4}
	fmt.Printf("main arr1 - %v", arr1)
	fmt.Println()

	arr2 := [4]int{1, 2, 3, 4}
	fmt.Printf("main arr2 - %v", arr2)
	fmt.Println()

	fnV(arr1)
	fnP(&arr2)

	fmt.Println()
	fmt.Println("Functions Run")
	fmt.Printf("main arr1 - %v", arr1)
	fmt.Println()
	fmt.Printf("main arr2 - %v", arr2)
	fmt.Println()
}
