package main

import (
	"fmt"
	"os"
)

func errCheck(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	/*
		C:\\Users\\kuron\\OneDrive\\바탕 화면\\golang_VSCode\\go_study\\src\\section11\\test_write.txt

		Open - 기존 파일 열기
		Close - 리소스 반환
		Read, ReadAt - 파일 읽기
		OS 권한 주의!
		예외 처리 중요!
		탐색 Seek 중요
		https://github.com/tealeg/xlsx
	*/

	//read File
	f, err := os.Open("C:\\Users\\kuron\\OneDrive\\바탕 화면\\golang_VSCode\\go_study\\src\\section11\\sample.txt")
	errCheck(err)

	fi, err := f.Stat() //file size를 위해 info 획득
	errCheck(err)

	fb := make([]byte, fi.Size()) //내용을 넣을 슬라이스 선언

	ct1, err := f.Read(fb)

	fmt.Println("파일 정보 출력 1 - ", fi)
	fmt.Println("파일 정보 출력 1 - ", fi.Name())
	fmt.Println("파일 정보 출력 1 - ", fi.Size())
	fmt.Println("- 파일 내용 - \n", string(fb))
	fmt.Println("읽은 길이 - ", ct1)
	fmt.Println("======================================")

	//read File with Seek(Offset)
	o1, err := f.Seek(20, 0) // whence - 0 : 처음위치, 1: 현재 위치(읽은 시점 이후), 2: 마지막 위치
	errCheck(err)

	fb2 := make([]byte, 20)
	ct2, err := f.Read(fb2)
	errCheck(err)

	fmt.Println("- 파일 내용 - \n", string(fb2))
	fmt.Println("읽은 길이 - ", ct2, o1)
	fmt.Println("======================================")

	o2, err := f.Seek(0, 0) // whence - 0 : 처음위치, 1: 현재 위치(읽은 시점 이후), 2: 마지막 위치
	errCheck(err)

	fb3 := make([]byte, 52)
	ct3, err := f.Read(fb3)
	errCheck(err)

	fmt.Println("- 파일 내용 - \n", string(fb3))
	fmt.Println("읽은 길이 - ", ct3, o2)
	fmt.Println("======================================")

	o3, err := f.Seek(0, 1) // whence - 0 : 처음위치, 1: 현재 위치(읽은 시점 이후), 2: 마지막 위치
	errCheck(err)

	fb4 := make([]byte, 56)
	ct4, err := f.Read(fb4)
	errCheck(err)

	fmt.Println("- 파일 내용 - \n", string(fb4))
	fmt.Println("읽은 길이 - ", ct4, o3)
	fmt.Println("======================================")

	defer f.Close()
}
