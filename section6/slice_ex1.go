package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3, 4, 5}
	fmt.Println(s1)
	fmt.Printf("%d %d", len(s1), cap(s1))
	fmt.Println()
	s1 = append(s1, 6, 7, 8)
	fmt.Println(s1)
	fmt.Printf("%d %d", len(s1), cap(s1))
	fmt.Println()

	s2 := make([]int, 5, 10)
	fmt.Println(s2)
	fmt.Printf("%d %d", len(s2), cap(s2))
	fmt.Println()
	//temp 에다가 2배 값 넣고 늘리고 옮기고 다시 s2에다가 넘겨주는 형식으로 용량을 넓힘

	s2 = append(s2, 1, 2, 3, 4, 5, 6, 7)
	fmt.Println(s2)
	fmt.Printf("%d %d", len(s2), cap(s2))
	fmt.Println()

	s3 := []int{8, 9, 10, 11, 12}
	fmt.Println(s3)
	fmt.Printf("%d %d", len(s3), cap(s3))
	fmt.Println()

	s4 := make([]int, 10, 20)
	s4 = append(s2, s3...)
	fmt.Println(s4)
	fmt.Printf("%d %d", len(s4), cap(s4))
	fmt.Println()

	s5 := make([]int, 0, 5)
	for i := 0; i < 21; i++ {
		s5 = append(s5, i)
		fmt.Printf("\nex5 : len : %d, cap : %d,  value: %v, \n", len(s5), cap(s5), s5)
	}

}
