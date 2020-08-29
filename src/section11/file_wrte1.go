package main

import (
	"fmt"
	"os"
)

//error 체크 방식1
func errCheck1(e error) {
	if e != nil {
		panic(e)
	}
}

func errCheck2(e error) {
	if e != nil {
		fmt.Println(e)
		return
	}
}

func main() {
	//파일 쓰기
	/*
		Create : 새 파일 작성 및 파일 열기
		Close : 리소스 반환
		Write, WriteString, WriteAt : 파일 쓰기
		각 OS의 권한들을 주의
		예외 처리 중요!
		http://golang.org/pkg/os
	*/

	f, err := os.Create("C:\\Users\\kuron\\OneDrive\\바탕 화면\\golang_VSCode\\go_study\\src\\section11\\test_write.txt")

	errCheck1(err)

	fmt.Println(f.Name())

	defer f.Close()

	//쓰기 byte
	s1 := []byte{115, 111, 109, 101, 115}
	n1, err := f.Write(s1) // 문자열 -> byte 슬라이스 형으로 변환 후 쓰기

	errCheck2(err)
	f.Sync()

	fmt.Printf("쓰기 작업(1) 완료 (%d bytes) \n", n1)

	//쓰기 string
	s2 := "\n가나다라마바사\n 아자아자아자젤\n"
	n2, err := f.WriteString(s2)

	errCheck1(err)
	f.Sync()

	fmt.Printf("쓰기 작업(2) 완료 (%d bytes) \n", n2)

	//쓰기 WriteAt
	s3 := "Test WriteAt \n"
	n3, err := f.WriteAt([]byte(s3), 55) //offset 조절하면서 테스트

	errCheck2(err)
	f.Sync()

	fmt.Printf("쓰기 작업(3) 완료 (%d bytes) \n", n3)

	//WriteString 바로 쓰기
	n4, err := f.WriteString("\nHello Golnag\n File wirte Test\n")

	errCheck1(err)

	fmt.Printf("쓰기 작업(4) 완료 (%d bytes) \n", n4)

}
