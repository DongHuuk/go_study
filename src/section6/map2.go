package main

import "fmt"

func main() {
	//Map조회 및 순회(Iterator)

	map1 := map[string]string{
		"daum":   "www.daum.co.kr",
		"naver":  "www.naver.com",
		"google": "www.google.co.kr",
	}

	fmt.Println("ex1 : ", map1["google"])
	fmt.Println("ex1 : ", map1["naver"])
	fmt.Println()

	for k, v := range map1 {
		fmt.Println("key = ", k)
		fmt.Println("value = ", v)
	}

}
