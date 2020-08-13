package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "Package fmt implements formatted I/O with functions analogous to C's printf and scanf. The format 'verbs' are derived from C's but are simpler." +
		"%v	the value in a default format" +
		"%#v	a Go-syntax representation of the value"

	str2 := "%T	a Go-syntax representation of the type of the value"

	fmt.Println("ex 1 : ", str1, str2)

	strSet := []string{} // 스트링 슬라이스
	strSet = append(strSet, str1)
	strSet = append(strSet, str2)

	fmt.Println("ex 1 : ", strSet)

	fmt.Println("ex 1 : ", strings.Join(strSet, ""))

}
