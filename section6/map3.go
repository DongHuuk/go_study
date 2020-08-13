package main

import "fmt"

func main() {
	map1 := map[string]string{
		"daum":   "www.daum.co.kr",
		"naver":  "www.naver.com",
		"google": "www.google.co.kr",
		"home":   "www.localhost:8080.co.kr",
	}

	fmt.Println("ex1 : ", map1)

	map1["home2"] = "www.localhost.com"

	fmt.Println("ex1 : ", map1)

	map1["home2"] = "www.localhost_test.com"

	fmt.Println("ex1 : ", map1)

	fmt.Println()

	delete(map1, "home2")

	fmt.Println("ex1 : ", map1)

	delete(map1, "naver")

	fmt.Println("ex1 : ", map1)

	fmt.Println()
}
