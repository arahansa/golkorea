/*
    author twitter : @mmcgrana
    https://gobyexample.com/number-parsing
	Go by Example: number-parsing

	문자열에서 숫자를 파싱하시는 것은 기초적이지만, 많은 프로그램에서 흔한 작업이다.
	여기에 어떻게 Go에서 그것을 하는지 나와있다.
*/

package main

// 내장된 패키지 strconv는 숫자파싱을 제공한다.
import "strconv"
import "fmt"

func main() {

	// `ParseFloat`와 함께 이 `64`는 얼마나 많은 비트수가 파싱될 것인지 말해준다.
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)

	// `ParseInt`에서 `0` 는 means infer the base from the string (번역요구)
	// 64는 결과가 64비트에 맞춰질 것을 요구한다.
	i, _ := strconv.ParseInt("123", 10, 64)
	fmt.Println(i)

	// ParseInt는 hex-형식숫자를 인식한다.
	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	//ParseUint 또한 사용가능하다.
	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)

	//`Atoi` 는 기초적인 10진수 파싱을 위한 편리한 함수이다.
	k, _ := strconv.Atoi("135")
	fmt.Println(k)

	// 나쁜 입력값에서 파싱함수는 에러를 리턴한다.
	_, e := strconv.Atoi("wat")
	fmt.Println(e)
}
