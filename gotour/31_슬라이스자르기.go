// 31_슬라이스자르기
package main

import "fmt"

func main() {
	p := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("p ==", p)
	fmt.Println("p[2:3] ==", p[2:3])

	// missing low index implies 0
	fmt.Println("p[:3] ==", p[:3])

	// missing high index implies len(s)
	fmt.Println("p[4:] ==", p[4:])
}

/*
슬라이스에서 슬라이스이름[시작:끝-1] 으로 재분할가능하다.
 예) 저기 위에서 슬라이스이름[2:3] 은 0, 1, 2에서 2번째것 즉 5만 나타냄.


 Written by arahansa
 https://www.facebook.com/groups/golangko/
모든 내용은 http://go-tour-kr.appspot.com/
*/
