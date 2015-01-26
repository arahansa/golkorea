/*
    author twitter : @mmcgrana
    https://gobyexample.com/collection-functions

	Go by Example: collection-functions

	우리는 종종 우리의 프로그램이 데이터콜렉션 위에서 작업을 수행하기를 원한다.
	주어진 술어(predicate)를 만족시키는 모든 아이템들을 selecting 하거나
	커스텀 함수를 이용하여 모든 item들을 새로운 콜렉션에 매핑하는 일같은 것들 말이다.

	어떤 언어에서는 제너릭 데이터구조와 알고리즘을 사용하는 것이 흔하다.
	Go는 제너릭을 지원하지 않는다.
	당신의 프로그램과 데이터 타입이 특별히 제너릭을 필요로 할때는 Go에서는 콜렉션함수를 제공하는 것이 흔하다.

	여기 문자열 슬라이스를 위한 콜렉션함수의 예제가 있다. 이 예제를 이용해서 당신 자신만의 함수를 만들 수 있다.
	어떤 경우에는 헬퍼함수를 생성하고 호출하는 것보다 코드를 직접적으로 조정하며 콜렉션을 처리하는 것이 제일 깔끔할 수도 있다.
*/

package main

import "strings"
import "fmt"

//  타겟문자열t 의 인덱스를 리턴하거나, 매칭되는게 없다면 -1을 리턴한다.
func Index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

// 만약 타겟문자열 t가 슬라이스에 있다면 true를 리턴한다.
func Include(vs []string, t string) bool {
	return Index(vs, t) >= 0
}

// 만약 슬라이스 안의 문자열중의 하나가 f(역주: 단어뜻은 술어? 인데 여기선 f func를 말하는듯)를 만족하면
// true를 리턴한다.
func Any(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

// 만약 슬라이스의 모든 문자열이 f술어를 만족하면 참을 리턴한다.
func All(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

// f를 만족시키는 슬라이스에서 모든 문자열을 가지고 있는 새로운 슬라이스를 리턴한다.
// 원문 참조
// Returns a new slice containing all strings in the
// slice that satisfy the predicate `f`.
func Filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

// 원래 슬라이스에있는 문자열마다 func f를 적용시킨 결과를 가지고 있는 새로운 슬라이스를 리턴한다.
func Map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

func main() {
	// 여기 우리가 테스트해보려는 다양한 콜렉션 함수가 있다.
	var strs = []string{"peach", "apple", "pear", "plum"}

	fmt.Println(Index(strs, "pear"))

	fmt.Println(Include(strs, "grape"))

	fmt.Println(Any(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))

	fmt.Println(All(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))

	fmt.Println(Filter(strs, func(v string) bool {
		return strings.Contains(v, "e")
	}))

	//위의 예제에서 우리는 익명함수를 사용했지만 당신은 또한 올바른 타입의 이름있는 함수를 사용할 수 있다.
	fmt.Println(Map(strs, strings.ToUpper))

}
