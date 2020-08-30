package main

import (
	"fmt"

	"github.com/tealeg/xlsx"
)

func cellVisitor(c *xlsx.Cell) error {
	value, err := c.FormattedValue()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("%s	\t", value)
	}
	return err
}

func rowVisitor(r *xlsx.Row) error {
	// fmt.Println("\n======")
	fmt.Println()
	return r.ForEachCell(cellVisitor)
}

func main() {
	//외부 패키지 설치
	//1. import 선언 후 폴더 이동 후 go get으로 설치
	//2. go get 패키지 주소 선언으로 설치

	fileName := "sample.xlsx"
	xFile, err := xlsx.OpenFile(fileName)

	//excel sheet가 index 0, 1, 2, 3 ...
	if err != nil {
		panic(err)
	}

	for shIndex, sheet := range xFile.Sheets {
		fmt.Println("index ", shIndex)
		fmt.Println("sheet ", sheet.Name)

		err = sheet.ForEachRow(rowVisitor)
		fmt.Println("Err = ", err)
	}

}
