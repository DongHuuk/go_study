package main

import (
	"fmt"
)

func receiveOnly(cnt int) <-chan int {
	sum := 0
	tot := make(chan int) //지역변수 chan
	go func() {
		for i := 0; i <= cnt; i++ {
			sum += i
		}
		tot <- sum // return 하면 그 return 값을 받은 변수는 수신 전용 chan이 된다.
		tot <- 777
		tot <- 777
		close(tot)
	}()
	return tot
}

func total(c <-chan int) <-chan int {
	tot := make(chan int)
	go func() {
		a := 0
		for i := range c {
			a += i
		}

		tot <- a
	}()
	return tot
}

func main() {
	c := receiveOnly(100) //수신 전용 chan
	output := total(c)

	fmt.Printf("%T\n", output)
	fmt.Println("ex1 : ", <-output)
}
