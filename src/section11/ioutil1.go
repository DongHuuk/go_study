package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func errCheck(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	/*
		파일 읽기 쓰기 -> ioutil 패키지 활용
		편리하고 직관적으로 파일을 쓰고 읽기 가능
		WriteFile(), ReadFile(), REadAll() 등 사용 가능
		https://golang.org/pkg/io/ioutil
	*/

	s := "Hello Golang!\nFile Wirte Test"

	//파일모드 (chmod, chown, chgrp) -> 퍼미션
	//Golang에서 사용이 가능하지만 Editer가 관리자 권한으로 동작해야지 동작한다. r-4, w-2, x-1, 소유자(user), 그룹(group), 기타 사용자(other)
	//https://golang.org/pkg/io/#FileMode

	err := ioutil.WriteFile("test_write1.txt", []byte(s), os.FileMode(0744)) //FileMode 에서 앞에 0을 붙여준다(그냥)
	errCheck(err)

	data, err := ioutil.ReadFile("test_write1.txt")
	errCheck(err)

	fmt.Println(data)
	fmt.Println("==================")
	fmt.Println(string(data))

	data2, err := os.Open("test_write1.txt")
	errCheck(err)

	dataInfo, err := data2.Stat()
	errCheck(err)

	d := make([]byte, dataInfo.Size())

	i, err := data2.Read(d)
	errCheck(err)

	fmt.Println(string(d))
	fmt.Println("read index - ", i)

	defer data2.Close()

}
