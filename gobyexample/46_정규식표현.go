/*
    author twitter : @mmcgrana
    https://gobyexample.com/regular-expressions
	Go by Example: regular-expressions

	Go는 내장된 정규식 표현을 지원하며, 여기 Go에서의 정규식관련작업 예제가 있다.
*/

package main

import "bytes"
import "fmt"
import "regexp"

func main() {

	// 이것은 문자열이 패턴에 맞는지 검사하는 것이다.
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	// 위에서 문자열패턴을 직접 사용했지만, 다른 정규식 작업을 위해
	// 최적화된 `Regexp`구조체를 `Compile`할 필요가 있을 것이다
	r, _ := regexp.Compile("p([a-z]+)ch")

	// 이 구조체에 많은 메소드들이 사용가능하다.
	//여기 우리가 전에서 본것과 비슷한 매치 테스트가 있다.
	fmt.Println(r.MatchString("peach"))

	// 이것은 정규식에 매칭되는 것을 찾아낸다.
	fmt.Println(r.FindString("peach punch"))

	// 이것은 첫번째 매칭되는 것을 찾을 뿐만 아니라, 매칭된 텍스트가 아닌 매칭된 것의 처음과 끝의 인덱스값을 리턴한다.
	fmt.Println(r.FindStringIndex("peach punch"))

	// `Submatch` 이 변종녀석은 정보를 포함하는데, 매칭되는 전체매칭과 그러한 전체 매칭에서의 하위매칭을 보여준다.
	// (역주: 정규식이 p로 시작하고 ch로 끝나는 것인데 먼저 peach를 찾고 peach에서 p, ch를 자른 ea를 리턴함)
	// 예제로, 이것은 p([a-z]+)ch 와 ([a-z]+)에 대한 정보 둘다를 리턴한다.
	fmt.Println(r.FindStringSubmatch("peach  punch"))

	// 비슷하게 이것은 매칭되는 것의 인덱스와 서브매칭되는 것의 인덱스를 리턴한다.
	fmt.Println(r.FindStringSubmatchIndex("peach punch"))

	// 이 'All' 변형은 입력에 대해서 첫번째값만이 아닌 모든 매칭값에 적용된다.
	// 예제에서 정규식에 맞는 모든 매칭값을 찾을 것이다.
	fmt.Println(r.FindAllString("peach punch pinch", -1))

	// 이러한 'All' 변형은 우리가 위에서 보았던 다른 함수들에도 적용가능하다.
	fmt.Println(r.FindAllStringSubmatchIndex(
		"peach punch pinch", -1))

	// 네거티브하지 않은 정수(양수)를 두번째 파라미터로 제공함으로써 매칭되는 것의 개수를 정할 수 있다.
	fmt.Println(r.FindAllString("peach punch pinch", 2))

	// 우리의 위의 예제에서 문자열 파라미터를 썼었고 `MatchString` 같은 이름을 썼었지만,
	// 우리는 또한 []byte 아규먼트를 주고 함수 이름에서 'String'을 얻을 수 있다.
	fmt.Println(r.Match([]byte("peach")))

	// 정규식표현과 함께 상수(constants)를 생성할 때, 우리는 'Compile'의 변종인 `MustCompile`을 사용할 수 있다.
	// Compile은 2개의 리턴값을 가지므로 상수에서 쓰일 수 없다.
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println(r)

	// `regexp` 패키지는 어떤 다른 값을 가지고, 문자열의 일부분을 바꾸는 데 사용할 수도 있다.
	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

	// `Func`변종은 매칭된 텍스트를 주어진 함수로 함께 바꿀 수 있게 해준다.
	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}
