package main

import (
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
	//csv file 쓰기 (csv file == 엑셀 파일)
	//패키지 저장소를 통해서 Excel 등 다양한 파일 형식 쓰기, 읽기 가능
	//Package Github address : https://github.com/tealeg/xlsx
	//bufid = buffer IO -> 파일이 용량이 클 경우 버퍼 사용 권장

	f, err := os.Create("C:\\Users\\kuron\\OneDrive\\바탕 화면\\golang_VSCode\\go_study\\src\\section11\\test_write.csv")
	errCheck(err)

	defer f.Close()

	//csv Writer 생성
	wr := csv.NewWriter(f)
	//wr := csv.NewWriter(bufio.NewWriter(file)) <- 버퍼를 추가해서 사용할 경우

	//csv Write
	wr.Write([]string{"Kim", "4.4"})
	wr.Write([]string{"Lee", "2.8"})
	wr.Write([]string{"Choi", "3.6"})
	wr.Write([]string{"Park", "1.8"})

	wr.Flush() //버퍼에 쌓아놨다가 파일에 쓰는 것

	fi, err := f.Stat()

	errCheck(err)

	fmt.Printf("CSV 쓰기 작업 후 파일 크기 (%d bytes)\n", fi.Size())
	fmt.Println("CSV FileName", fi.Name())
	fmt.Println("CSV Mode", fi.Mode()) // 권한
	fmt.Println("CSV is Directory? ", fi.IsDir())
	fmt.Println("CSV Sys ", fi.Sys())
	fmt.Println("CSV ModeTime ", fi.ModTime())

}
