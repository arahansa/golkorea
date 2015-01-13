// 61_연습Rot13Reader
package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func main() {
	s := strings.NewReader(
		"Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

/*
어떤 식으로든 스트림을 수정하여 다른 io.Reader 를 감싸는 io.Reader 는 흔한 패턴입니다.

예컨대, gzip.NewReader 함수는 io.Reader (gzip으로 압축된 데이터의 스트림) 를 가지고,
 io.Reader (압축 해제된 데이터의 스트림) 를 구현한 `*gzip.Reader`를 반환합니다.

ROT13 치환 암호화를 모든 알파벳 문자에 적용함으로써 스트림을 수정하며
io.Reader 를 구현하고 io.Reader 로 부터 읽는 rot13Reader 를 구현하십시오.
rot13Reader 타입은 당신을 위해 제공됩니다. 이 타입의 Read 함수를 구현함으로써 io.Reader 을 만들어 보십시오.

Written by arahansa
https://www.facebook.com/groups/golangko/
모든 내용은 http://go-tour-kr.appspot.com/
*/
