// 41_연습_맵
package main

import (
	"code.google.com/p/go-tour/wc"
)

func WordCount(s string) map[string]int {
	return map[string]int{"x": 1}
}

func main() {
	wc.Test(WordCount)
}

/*
WordCount 함수를 구현합니다. 이 함수는 s라는 문자열 내에서 각각의 "단어"의 등장 횟수를 나타내는 맵을 반환해야 합니다.
wc.Test 함수는 주어진 함수를 이용하여 테스트를 실행한 뒤에 그 성공 여부를 출력해 줍니다.

아마도 다음 링크 (strings.Fields)의 내용이 도움이 될 것입니다.

Written by arahansa
https://www.facebook.com/groups/golangko/
모든 내용은 http://go-tour-kr.appspot.com/
*/
