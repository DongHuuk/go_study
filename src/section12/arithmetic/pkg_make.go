//사용자 패키지 작성 및 문서화 에제
package main

import (
	oper "arithmetic" //alias 사용 = oper
	"fmt"
)

//main은 godoc에 포함되지 않는다.
func main() {
	/*
		기준은 GOPATH/src
		패키지 폴더명(디텍토리명) 명확하게 지정
		패키지 파일의 package 이름으로 사용한다. -> alias 사용 자유
		package main을 제외하고 package 문서에 등록
		GOROOT pkg 검색 -> GOPATH의 pkg(src/pkg) 검색
		go install 명령어 실행의 경우, GOPATH/pkg에 등록(문서화 자동)
		godoc -http=:6060(임의의 포트) -> pkg 이동 -> 본인 패키지 메서드 및 주석 확인(패키지, 타입, 메서드)
	*/

	nums := oper.Numbers{X: 100, Y: 500}

	fmt.Println("ex1 : ", nums.SumNumber())
	fmt.Println("ex1 : ", nums.MinusNumber())
	fmt.Println("ex1 : ", nums.MultiNumber())
	fmt.Println("ex1 : ", nums.DivideNumber())
	fmt.Println("ex1 : ", nums.Plus())
	fmt.Println("ex1 : ", nums.Minus())

}
