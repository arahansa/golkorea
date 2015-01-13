// 54_인터페이스의암시적충족
package main

import (
	"fmt"
	"os"
)

type Reader interface {
	Read(b []byte) (n int, err error)
}

type Writer interface {
	Write(b []byte) (n int, err error)
}

type ReadWriter interface {
	Reader
	Writer
}

func main() {
	var w Writer

	// os.Stdout implements Writer
	w = os.Stdout

	fmt.Fprintf(w, "hello, writer\n")
}

/*
타입이 인터페이스의 메소드들을 구현하면 인터페이스를 구현한 게 됩니다.
이를 위해 명시적으로 선언할 게 없습니다.
암시적 인터페이스는 인터페이스를 정의한 패키지로 부터
구현 패키지를 분리(decouple)해 줍니다. 다른 의존성 또한 없음은 물론입니다.
이 특징은 상세하게 인터페이스를 정의하게 독려합니다. 모든 구현을 찾아
새 인터페이스 이름으로 태그할 필요가 없기 때문입니다.
패키지 io에 Reader 와 Writer 가 정의되어 있습니다. 따로 정의할 필요가 없습니다.


Written by arahansa
https://www.facebook.com/groups/golangko/
모든 내용은 http://go-tour-kr.appspot.com/
*/
