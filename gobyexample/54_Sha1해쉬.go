/*
    author twitter : @mmcgrana
    https://gobyexample.com/sha1-hashes
	Go by Example: sha1-hashes

	SHA1 hashes는 바이너리나 텍스트 blob를 위한 짧은 정체들을(identities) 계산하는데 자주 사용된다.
	예를들면 git 리비젼 컨트롤시스템은 Sha1을  파일과 디렉토리의 버젼확인을 위해 널리 사용을 한다.
	여기 어떻게 GO에서 SHA1 해쉬를 계산하는지 나와있다.

*/
package main

// Go는 많은 crypto/* 패키지에서 몇몇의 해쉬 함수를 구현한다.
import "crypto/sha1"
import "fmt"

func main() {
	s := "sha1 this string"

	// 해쉬를 만드는 패턴은 sha1.New() , sha1.Write(bytes) 그리고 sha1.Sum([]bytes{}) 이다.
	//여기 새로운 해쉬로 시작하는 방법이 있다.
	h := sha1.New()

	// Write는 bytes를 기대한다(Expects) 만약 당신이 문자열 s 를 가지고 있다면,
	// 그것을 byte로 강제하기 위해 []byte(s)를 사용하라
	h.Write([]byte(s))

	// 이것은 확정된 해쉬값을 바이트슬라이스로 받는다
	// Sum의 아규먼트는 존재하는 바이트 슬라이스에서 더 추가하기 위해 사용될 수 있다.(보통 필요하지 않다)
	bs := h.Sum(nil)

	// 깃커밋의 예를 들자면 SHA1 값은 종종 hex값으로 출력된다.
	// %x 형식을 사용하여서 해쉬값을 hex문자열로 바꿔보자.
	fmt.Println(s)
	fmt.Printf("%x\n", bs)
}
