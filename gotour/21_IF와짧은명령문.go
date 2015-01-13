// 21_IF와짧은명령문
//IF실행전 간단한 변수선언가능하다.
package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}

/*
 if문 앞에 짧은 문장을 넣을 수 있다.  이부분 =>  v := math.Pow(x, n);
 if 블록 안에서만 사용가능하다.

 Written by arahansa
 https://www.facebook.com/groups/golangko/
모든 내용은 http://go-tour-kr.appspot.com/
*/
