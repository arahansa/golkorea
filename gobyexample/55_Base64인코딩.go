/*
    author twitter : @mmcgrana
   	https://gobyexample.com/base64-encoding
	Go by Example: base64-encoding

	Go는 base64를 위한 내장형 지원을 제공한다.
*/

package main

// 이 구문은 `encoding/base64`패키지를 디폴트값(base64)가 아닌 b64라는 이름으로 임포트 한다.
// 그것은 우리에게 공간을 좀 줄여주게 할 것이다.
import b64 "encoding/base64"
import "fmt"

func main() {

	// 여기 우리가 인코드/디코드할 문자열이 있다
	data := "abc123!?$*&()'-=@~"

	// Go는 표준과 URL호환 base64둘다 지원한다.
	// 여기 어떻게 표준인코더를 사용하여서 어떻게 인코드하는지 나와있다.
	// Encoder는 []byte를 필요로 하기 때문에 우리는 문자열을 바이트 타입으로 형변환할 것이다.
	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)

	// 디코딩은 아마 에러를 리턴할 것이고, 그 에러를 가지고서
	// 당신이 이미 잘 적혀진 입력값을 모르고 있는지 체크할 수 있다.
	// Decoding may return an error, which you can check
	// if you don't already know the input to be well-formed.
	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	fmt.Println("sdec:", string(sDec))
	fmt.Println()

	// 이것은 base64형식의 URL호환인코딩디코딩이다.
	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)
	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))
}
