/*
    author twitter : @mmcgrana
   	https://gobyexample.com/command-line-flags
	Go by Example: command-line-flags

	커맨드라인플래그는 커맨드라인프로그램을 위한 옵션을 명시하는 흔한 방법이다.
	예를 들자면 wc -l 에서 -l 은 커맨드라인플래그이다.
*/

package main

// Go는 flag패키지를 제공하여, 기본적인 커맨드라인 플래그 파싱을 지원한다.
// 우리는 이 패키지를 사용하여 우리의 커맨드라인프로그램을 구현해보고자 한다.
import "flag"
import "fmt"

func main() {

	// 기본 플래그 선언은 문자열, 정수, 불린형으로 선언 가능하다.
	//여기 문자열 플래그 word를 기본값인 foo와 짧은 설명과 함께 선언했다.
	// 이 flag.String 함수는 string포인터를 리턴한다. (String값이 아니다)
	// 우리는 이 포인터를 어떻게 쓸 것인지 밑에서 볼 것이다.
	wordPtr := flag.String("word", "foo", "a string")

	// 이것은 numb와 fork 플래그를 선언한다.
	// word플래그와 비슷한 접근을 한다.
	numbPtr := flag.Int("numb", 42, "an int")
	boolPtr := flag.Bool("fork", false, "a bool")

	// 다른 곳에서 선언하여 프로그램에 이미 존재하는 변수를 사용하는 옵션을 선언하는 것도 가능하다.
	// flag선언 함수에 포인터를 전달해야 한다는 것을 기억하자
	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	// 한번 모든 플래그들이 선언되면 flag.Parse() 를 호출하여 커맨드라인파싱을 실행한다.
	flag.Parse()

	// 여기 우리는 파싱된 옵션들과 그것에 따르는 아규먼트들을 출력한다.
	//우리는 포인터를 디레퍼런스 해야한다는 것을 기억하자
	// 예를 들자면 *wordPtr은 실제 옵션값을 얻는다.
	// Here we'll just dump out the parsed options and
	// any trailing positional arguments. Note that we
	// need to dereference the points with e.g. `*wordPtr`
	// to get the actual option values.
	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *boolPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())
}
