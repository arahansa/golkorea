// 34_레인지
package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}

/*
for 반복문에서 range 를 사용하면 슬라이스나 맵을 순회(iterates)할 수 있다.

 Written by arahansa
 https://www.facebook.com/groups/golangko/
모든 내용은 http://go-tour-kr.appspot.com/
*/
