// RETURN 시 반환 안 정해주게도함.
package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
}

/*
 내용 : 함수는 매개변수를 취하는데, 리턴할 이름을 정해주면 알아서 반환한다..이 말임.
 잘 볼 곳 : return 만 적었는데 반환값이 x, y int 이므로 x,y 를 알아서 리턴함.

 Written by arahansa
 https://www.facebook.com/groups/golangko/
모든 내용은 http://go-tour-kr.appspot.com/
*/
