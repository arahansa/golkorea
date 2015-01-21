// strconv 패키지는 전환을 구현한다. 문자열로의 아니면, 문자열로부터 기본데이터타입의 표현으로~
//Package strconv implements conversions to and from string representations of basic data types.

package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Hello world")
	//Atoi 는 문자열=> 숫자로 파싱해준다. 두번째 값은 에러리턴.
	fmt.Println(strconv.Atoi("27"))
	//Itoa 는 숫자를 문자열로 파싱
	fmt.Println("숫자=>문자는 Itoa :", strconv.Itoa(27))
	// 파라미터는 (i int64, base int) 형으로, base는 2(포함)와 36(포함)사이의 값
	fmt.Println(strconv.FormatInt(11, 36))
}
