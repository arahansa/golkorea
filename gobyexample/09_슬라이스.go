/*
	author twitter : @mmcgrana
	https://gobyexample.com/slices

	Go by Example: Slices
*/
package main

import "fmt"

func main() {

	// 배열과 달리, 오직 슬라이스가 가지고 있는 요소들의 타입을 속성으로 가진다.
	// 이 말은 배열은 요소타입, 길이를 가지는데 슬라이스는 요소타입만 가진다는 얘기.(역주, 가변길이)
	// 가변길이이기 때문에 이따가 더 추가하는 것이 나온다. 한번 보자!!

	// 길이가 0이 아닌 슬라이스를 만들기 위해,
	// 내장 기능인 make 를 이용해봅니다. 이것은 문자열, 크기3을 가진 슬라이스를 만듭니다.
	// zero-value 로 초기화됩니다.
	s := make([]string, 3)
	fmt.Println("emp:", s)

	// 배열처럼 요소들의 값을 저장하고 불러올 수 있습니다.
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	//len 은 길이 리턴하시는 것~ 기억하시죠~
	fmt.Println("len:", len(s))

	// 배열보다 더 많은 기능중의 하나는 append 이다.
	// append는 하나 혹은 그 이상의 새로운 값을 포함하는 슬라이스를 리턴한다.
	// append 시에 리턴 값을 받아야 한다는 것을 잊지마세요!
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	// 슬라이스는 또한 `copy` 될 수 있다.여기에 빈 슬라이스 c를 만드는데,
	//길이는 s의 길이로 만들어서 s 의 내용을 c로 복사하는 내용이 나온다.
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	// 슬라이스는 연산자를 제공하는데 슬라이스[low:high]이다.
	// 이것은  `s[2]`, `s[3]`, and `s[4]` 를 얻는다.
	l := s[2:5]
	fmt.Println("sl1:", l)

	// 이것은 s[5]이전것(비포함)까지를 보여준다.
	l = s[:5]
	fmt.Println("sl2:", l)

	//  `s[2]` 부터(포함) 보여준다.
	l = s[2:]
	fmt.Println("sl3:", l)

	// 한줄로 또한 슬라이스를 선언&초기화할 수 있다. (배열과는 다르게 [숫자]가 안 들어감을 보자.역주)
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	// 슬라이스또한 다차원이 될 수 있다. 다차원 배열과는 다르게 내부의
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
	//이곳에 더 많은 자료가 있다고 한다. http://blog.golang.org/go-slices-usage-and-internals
}
