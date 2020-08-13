package main

import "fmt"

func main() {

	var str1 = "Golang"
	var str2 = "World"

	fmt.Println("ex1 : ", str1[0:2], str1[0])
	fmt.Println("ex1 : ", str2[3:], str2[0])
	fmt.Println("ex1 : ", str2[:3])
	fmt.Println("ex2 : ", str1+str2)
}
