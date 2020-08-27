package main

import (
	"fmt"
)

func main() {
	//chan 수신된 값과 수신이 정상적으로 이루어졌는가에 대한 값을 받을 수 있다

	ch := make(chan string)

	go func() {
		for i := 0; i < 3; i++ {
			ch <- "Good"
		}
	}()

	val1, ok1 := <-ch
	fmt.Println("ex1 : ", val1, ok1)
	val2, ok2 := <-ch
	fmt.Println("ex1 : ", val2, ok2)
	val3, ok3 := <-ch
	fmt.Println("ex1 : ", val3, ok3)

	close(ch)
	val4, ok4 := <-ch
	fmt.Println("ex1 : ", val4, ok4)

}
