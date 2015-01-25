/*
    author twitter : @mmcgrana
    https://gobyexample.com/sorting

	Go by Example: sorting

	Go의 sort패키지는 내장형이거나 사용자가정의한 타입에 관하여 정렬을 구현한다.
	내장형에 관한 정렬을 먼저 알아보겠다
*/

package main

import "fmt"
import "sort"

func main() {

	// Sort메소드는 내장타입에 대하여 구체적이다.
	// 여기 문자열에 관한 예제가 있다. 정렬은 제자리에서 한다는 것을 기억하자(sort. 이렇게~)
	// 그래서 주어진 슬라이스를 변경하지, 새로운 것을 리턴하지 않는다.
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	// int형을 정렬하는 예제다.
	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Ints:   ", ints)

	// 우리는 또한 슬라이스가 이미 정렬되었는지 체크할 수 있다.
	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted: ", s)

	//우리의 프로그램을 실행하면 정렬된 문자열과 숫자 슬라이스, 그리고 참값을
	// AreSorted테스트의 결과로 나타낼 것이다.
}
