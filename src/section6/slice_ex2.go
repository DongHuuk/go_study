package main

import (
	"fmt"
	"sort"
)

func main() {
	//슬라이스 추출 및 정렬
	s1 := make([]int, 0, 5)
	s1 = append(s1, 1, 2, 3, 4, 5)

	fmt.Println(s1[0:2])
	fmt.Println(s1[3:])
	fmt.Println(s1[:3])
	fmt.Println(s1[:])

	//sort pakage https://golang.org/pkg/sort
	s2 := []int{6, 5, 2, 8, 9, 1, 0}
	s3 := []string{"a", "c", "d", "g", "z", "y", "b"}
	fmt.Println("ex2: ", sort.IntsAreSorted(s2))
	sort.Ints(s2)
	fmt.Println("ex2: ", s2)
	fmt.Println("ex2: ", sort.IntsAreSorted(s2))
	fmt.Println()
	fmt.Println("ex3: ", sort.StringsAreSorted(s3))
	sort.Strings(s3)
	fmt.Println("ex3: ", s3)
	fmt.Println("ex3: ", sort.StringsAreSorted(s3))
}
