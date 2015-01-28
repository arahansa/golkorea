/*
    author twitter : @mmcgrana
    https://gobyexample.com/string-functions

	Go by Example: string-functions

	표준라이브러리의 string 패키지는 많은 유용한 문자열관련함수를 제공하는데
	여기에 패키지에 대한 이해를 도울만한 예제를 제공한다.
*/

package main

import s "strings"
import "fmt"

// 우리는 fmt.Println을 짧은 별명을 지어서 밑에와 같이 많이 사용할 수 있다
var p = fmt.Println

func main() {

	// 여기 문자열에서 사용가능한 함수예제가 있다.
	// string패키지에서 온 이러한 모든 함수는 string객체 자신의 메소드가 아니란 것을 주의하자.
	// 이것은 질문(밑의 HasPrefix등등을 말하는 듯 : 역주)에서
	// 문자열을 함수의 첫번째 아규먼트로 전달해야 된다는 것을 의미한다.

	p("Contains:  ", s.Contains("test", "es"))
	p("Count:     ", s.Count("test", "t"))
	p("HasPrefix: ", s.HasPrefix("test", "te"))
	p("HasSuffix: ", s.HasSuffix("test", "st"))
	p("Index:     ", s.Index("test", "e"))
	p("Join:      ", s.Join([]string{"a", "b"}, "-"))
	p("Repeat:    ", s.Repeat("a", 5))
	p("Replace:   ", s.Replace("foo", "o", "0", -1))
	p("Replace:   ", s.Replace("foo", "o", "0", 1))
	p("Split:     ", s.Split("a-b-c-d-e", "-"))
	p("ToLower:   ", s.ToLower("TEST"))
	p("ToUpper:   ", s.ToUpper("test"))
	p()

	//당신은 좀더 많은 함수를 docs에서 찾을 수 있다. (http://golang.org/pkg/strings/)

	// string의 부분은 아니지만 언급할만한 가치가 있는 메커니즘은 string의 길이를 구하는 것과
	// index로 문자를 얻는 것이다.

	p("Len: ", len("hello"))
	p("Char:", "hello"[1])
}
