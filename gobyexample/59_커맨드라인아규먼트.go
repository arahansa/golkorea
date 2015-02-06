/*
    author twitter : @mmcgrana
   	https://gobyexample.com/command-line-arguments
	Go by Example: command-line-arguments

	커맨드라인 아규먼트는 프로그램 실행에서 파라미터를 쓰는 흔한 방법이다.
	예를 들어서 go run hello.go 에서 run과 hello.go 는 Go프로그램의 아규먼트들이다.
*/

package main

import "os"
import "fmt"

func main() {

	// os.Args는 커맨드라인 아규먼트에 접근을 제공한다.
	// 이 슬라이스의 첫번째 값은 프로그램에 path란 것을 기억하자.
	// os.Args[1:]은 프로그램의 아규먼트들을 가지고 있다.
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	// 당신은 평범한 인덱싱으로 독립적인 아규먼트를 얻을 수 있다.
	arg := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)

	// 커맨드 라인 아규먼트를 실험하기 위해서, 먼저 go build로 바이너리파일을 만들 것이 좋다.
}
