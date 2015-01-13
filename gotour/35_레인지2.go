// 35_레인지2
package main

import "fmt"

func main() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i)
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}

/*
_ 를 이용해서 인덱스(index)나 값(value)를 무시할 수 있습니다.
만약 인덱스만 필요하다면 “ `, value` ” 부분을 다 제거하면 됩니다.

for i, value := range pow {
	pow[i] = 1 << uint(i)
}
를
for i := range pow {
	pow[i] = 1 << uint(i)
}
로바꾸기 가능

 Written by arahansa
 https://www.facebook.com/groups/golangko/
모든 내용은 http://go-tour-kr.appspot.com/
*/
