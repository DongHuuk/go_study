package main

import (
	"bufio"
	"fmt"
	"os"
	"io/ioutil"
)

func errCheck(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	//버퍼 사용 -> bufio
	/*
		ioutil, bufio 등은 io.Reader, ioWriter 인터페이스를 구현함 즉,Writer, Read 메서드 사용 가능
		======================
		type Reader interface{
			Read(p []byte) (n int, err error)
		}

		type Writer interface{
			write(p []byte) (n int, err error)
		}
		======================
		= bufio의 NewReader, NewWriter를 통해서 객체 반환 후 메서드 사용 가능
	*/

	file, err := os.OpenFile("test_write2.txt", os.O_CREATE|os.O_RDWR, os.FileMode(0744))
	errCheck(err)

	// wt := bufio.NewWriter(file) //Writer 반환(버퍼 사용)
	wt := bufio.NewWriterSize(file, 1048576) // buffer size - 1MB setting
	wt.WriteString("Test Write With Buffer")
	wt.Write([]byte("\nTest Write With Buffer"))

	fmt.Printf("사용한 Buffer Size (%d bytes)\n", wt.Buffered())
	fmt.Printf("남은 Buffer Size (%d bytes)\n", wt.Available())
	fmt.Printf("전체 Buffer Size (%d bytes)\n", wt.Size())
	wt.Flush()
	fmt.Println("================================")

	rt := bufio.NewReader(file)
	fileInfo, err := file.Stat()
	errCheck(err)

	body := make([]byte, fileInfo.Size())

	file.Seek(0, os.SEEK_SET)
	data, _ := rt.Read(body)

	fmt.Println(data)
	fmt.Println(string(body))

	fmt.Printf("전체 Buffer size (%b bytes)", rt.Size())

	defer file.Close()	

	

}
