/*
    author twitter : @mmcgrana
   	https://gobyexample.com/line-filters
	Go by Example: line-filters

	line 필터는 프로그램의 보통 타입으로, stdin 에서 입력을 읽고 처리하며,
	나온 결과들을 stdout에 출력한다.
	grep과 sed는 line filter이다.

	여기 Go에서 line필터 예제가 있으며, 모든 입력글자들을 대문자로 쓴다.
	당신은 이러한 패턴을 당신 자신의 GO 라인필터에 쓸 수 있다.

	애매한 번역은 원문을 남긴다. 역주
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	// 버퍼되지 않은 os.Stdin을 버퍼된 스캐너로 포장하는 것은 우리에게 편리한 Scan메소드를 제공하며
	// 스캐너를 다음 토큰으로 진행시킨다. 기본 스캐너에서 다음토큰은 다음줄이다.
	// Wrapping the unbuffered `os.Stdin` with a buffered
	// scanner gives us a convenient `Scan` method that
	// advances the scanner to the next token; which is
	// the next line in the default scanner.
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		// 'Text'는 현재 토큰을 리턴한다. 여기 입력에서 받은 다음 토큰이 있다.
		// `Text` returns the current token, here the next line,
		// from the input.
		ucl := strings.ToUpper(scanner.Text())

		//대문자된 줄을 출력한다.
		// Write out the uppercased line.
		fmt.Println(ucl)
	}

	// Scan도중에 에러가 있는지 검사한다.
	// 파일의 끝이 기대되며,Scan에 의해 에러로 보고되지 않는다.
	// Check for errors during `Scan`. End of file is
	// expected and not reported by `Scan` as an error.
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
