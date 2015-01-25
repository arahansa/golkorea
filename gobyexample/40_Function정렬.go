/*
    author twitter : @mmcgrana
    https://gobyexample.com/sorting-by-functions

	Go by Example: Sorting by Functions

	때떄로 우리는 정렬된 순서와는 다른 것으로 콜렉션을 정렬하고 싶다.
	예를 들자면 우리가 문자열을 그들의 알파벳이 아닌 길이로 정렬하고 싶다고 하자.
	여기 Go에서 커스텀 Sort의 예제가 있다.
*/

package main

import "sort"
import "fmt"

// Go에서 커스텀펑션으로 정렬하기위해 우리는 그에 맞는 타입이 필요하다.
// 단지 []string 타입의 다른 이름일 뿐인 ByLength라는 타입을 만들었다.
type ByLength []string

// 우리는 sort.Interface- Len, Less, Swap 을 구현했다.  - 우리의 타입에서
// 우리는 sort 패키지의 일반적인 Sort 함수를 사용할 수 있을 것이다.
// Len과 Swap은 일반적으로 비슷하고 Less는 실제 커스텀 정렬 로직을 가지고 있다.
// (Len` and `Swap` will usually be similar across types : 참고)
// 우리의 경우, 우리는 길이에 따라 정렬하기를 원하므로 우리는 len(s[i])와 len(s[j])를 여기선 사용할 것이다.
func (s ByLength) Len() int {
	return len(s)
}
func (s ByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ByLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

// 준비된 이것들과 함께, 원본 'fruits'를 ByLength로 형변환하면서 우리의 커스텀정렬을 해볼 것이다.
// 그리고 타입슬라이스에 sort.Sort를 이용할 것이다.
func main() {
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(ByLength(fruits))
	fmt.Println(fruits)
}

// 우리가 의도한 대로 우리의 프로그램은 정렬된 리스트를 보여줄 것이다.
// 커스텀타입을 생성하는 이와같은 패턴을 따르면서, 저 타입에 세가지 Interface메소드를 구현하고
// 저 커스텀타입콜렉션에  sort.Sort 를 호출하여서 우리는 임의의 함수로 Go slice를 정렬할 수 있다
