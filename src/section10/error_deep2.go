package main

import (
	"fmt"
	"math"
)

func number(f, i float64) (float64, error) {
	if f == 0 {
		return 0, fmt.Errorf("(%g)은 사용 불가", f) // %g float 표시
	}

	if i == 0 {
		return 1, fmt.Errorf("(%g)은 사용 불가", i)
	}

	return math.Pow(f, i), nil
}

func main() {
	//Core Example

	if f, err := number(8, 3); err != nil {
		fmt.Printf("Error Message : %s\n", err)
	} else {
		fmt.Println("ex1 : ", f)
	}

	if f, err := number(0, 3); err != nil {
		fmt.Printf("Error Message : %s\n", err)
	} else {
		fmt.Println("ex1 : ", f)
	}

	if f, err := number(5, 0); err != nil {
		fmt.Printf("Error Message : %s\n", err)
	} else {
		fmt.Println("ex1 : ", f)
	}
}
