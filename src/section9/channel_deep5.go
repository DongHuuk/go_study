package main

import (
	"fmt"
	"time"
)

func main() {

	ch1 := make(chan int)
	ch2 := make(chan string)

	go func() {
		for {
			num := <-ch1 // 값 수신용
			fmt.Println("ch1 : ", num)
			time.Sleep(250 * time.Millisecond)
		}
	}()

	go func() {
		for {
			ch2 <- "Golang Hi!" // 값 송신
			time.Sleep(500 * time.Millisecond)
		}
	}()

	go func() {
		for {
			select {
			case ch1 <- 777: // 송신
			case str := <-ch2: // 수신
				fmt.Println("ch2 : ", str)

			}
		}
	}()
	
	time.Sleep(5 * time.Second)
}
