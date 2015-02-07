/*
    author twitter : @mmcgrana
   	https://gobyexample.com/exit
	Go by Example: exit

	os.Exit는 주어진 상황(status)과 함께 즉시 종료된다
*/

package main

import "fmt"
import "os"

func main() {

	// 'defer'는 os.Exit를 사용할 때 실행되지 않을 것이다.
	//그러므로 이 'fmt.Println'은 절대 호출되지 않을 것이다.
	defer fmt.Println("!")

	// status3과 함께 종료된다.
	os.Exit(3)
}

// C와는 달리 Go는 main에서 exit status를 가리키기 위해 정수리턴값을 사용하지 않는다.
// 만약 제로값이 아닌 status를 사용하면서 종료 하고 싶다면 os.Exit를 사용해야 한다.
