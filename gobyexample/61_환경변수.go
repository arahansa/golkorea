/*
    author twitter : @mmcgrana
   	https://gobyexample.com/environment-variables
	Go by Example: environment-variables

	환경변수[Environment variables]는 설정정보를 유닉스프로그램에 전달하기 위한 일반적인 메커니즘이다.
	어떻게 환경변수를 설정하고, 값을 얻고 리스트화하는지 알아보자
*/

package main

import "os"
import "strings"
import "fmt"

func main() {

	// 키밸류 페어를 설정하기 위해, os.Setenv를 이용하자.
	// 키값에 대한 값을 얻기 위해, os.Getenv를 이용하다.
	// 만약 환경에 맞는 키가 없으면 빈 문자열을 리턴할 것이다.
	os.Setenv("FOO", "1")
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))

	// 환경의 모든 키밸류 쌍을 리스트출력하기 위해 os.Environ을 사용하자.D
	//이것은 문자열의 슬라이스를 KEY=value 폼으로 리턴한다.
	// strings.Split으로 그것들을 키와 밸류로 나눌 수 잇다. 여기 모든 키를 출력하는 예제가 있다.
	fmt.Println()
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		fmt.Println(pair[0])
	}
}
