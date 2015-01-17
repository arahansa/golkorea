/*
	author twitter : @mmcgrana
	https://gobyexample.com/range

	Go by Example: range

    _range_ 는 다양한 데이터구조속에서 요소들을 반복한다.
	우리가 이미 배운 데이터구조에서 어떻게 range 를 사용하는지 알아봅시다.
*/
package main

import "fmt"

func main() {

	// 슬라이스에서 숫자들을 어떻게 합치는지 보여줍니다.
	// 배열도 이와 같이 작동합니다.
	// 역주) _, num 에서 _은 증가값이고 num은 데이터구조의 요소들입니다.
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	// 배열과 슬라이스에서의 `range` 는 각각의 요소에, 인덱스(순번)과 값을 제공하며
	//인덱스값이 필요하지 않을 때는 _로 생략할 수 있다.
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	// 맵에서의 range는 키/값 을 돌려줍니다.
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// 문자열에서의 `range` 는 유니코드 값에서 돌아간다.(역주: 다음에 rune 이 나오는데 뭔지 모르겠음)
	// 첫번째 값은 시작하는 the rune의 byte index고
	// 두번째값은 the `rune` 그 자신이다.
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
