/*
	author twitter : @mmcgrana
	https://gobyexample.com/functions

	Go by Example: Multiple Return Values
	
	Go는 다중리턴값을 지원합니다. 
	이러한 특징은 Go 에서 자연스럽게 사용되며, 예를 들자면 
	함수로부터 결과값과 에러 둘다 받을 때 사용됩니다. 
*/
package main

import "fmt"

	// 이 함수에서 (int, int) 는 두개의 int형을 리턴한다는 것을 의미합니다. 
func vals() (int, int) {
    return 3, 7
}

func main() {

    //여기, 다중리턴호출로부터 두개의 다른 값을 받아봅니다. 
    a, b := vals()
    fmt.Println(a)
    fmt.Println(b)

    
    //만약 다중리턴 에서 필요하지 않은게 있으면 그것은 _ 로 처리하면 된다. 
    _, c := vals()
    fmt.Println(c)
}

// todo: named return parameters
// todo: naked returns
