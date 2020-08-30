package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
)

func errCheck(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := os.Open("C:\\Users\\kuron\\OneDrive\\바탕 화면\\golang_VSCode\\go_study\\src\\section11\\sample.csv")

	errCheck(err)

	defer file.Close()

	rr := csv.NewReader(bufio.NewReader(file))

	row, err := rr.Read()
	errCheck(err)

	fmt.Println("CSV Example")
	fmt.Println(row)
	fmt.Println(row[3])
	fmt.Println(row[3:8])
	fmt.Println("============================")

	row2, err2 := rr.Read()
	errCheck(err2)

	fmt.Println("CSV Example Next Line")
	fmt.Println(row2)
	fmt.Println(row2[3])
	fmt.Println(row2[3:8])
	fmt.Println("============================")

	//ReadAll

	all, err := rr.ReadAll()
	errCheck(err)

	fmt.Println("CSV Example All Rows")
	fmt.Println(all[6])

	for i, row := range all[7] {
		fmt.Println("index - ", i)
		fmt.Println("value - ", row)
	}

}
