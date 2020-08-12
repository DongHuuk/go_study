package main

import "fmt"

func main() {
	i := 7
	var p1 *int = new(int)
	p1 = &i

	fmt.Println(i)
	fmt.Println(p1)
	fmt.Println(*p1)
}
