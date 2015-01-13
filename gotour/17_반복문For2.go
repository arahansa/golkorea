// 17_For2
package main

import "fmt"

func main() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

}

/*

전후처리를 제외하고 조건문만 표현할 수도 있다.
무한루프는 그냥 for {} 로 표시.

 Written by arahansa
 https://www.facebook.com/groups/golangko/
모든 내용은 http://go-tour-kr.appspot.com/
*/
