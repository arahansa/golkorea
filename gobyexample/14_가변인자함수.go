/*
	author twitter : @mmcgrana
	https://gobyexample.com/variadic-functions

	Go by Example: Variadic Functions
	
 [_Variadic functions_](http://en.wikipedia.org/wiki/Variadic_function)
 가변인자 함수는 여러개의 파라미터를 받을 수 있는 함수다. 
 예를 들자면 fmt.Println이 가변인자함수 
*/
package main

import "fmt"

// 임의의 숫자 nums 들을 int 로 받는 함수 
func sum(nums ...int) {
    fmt.Print(nums, " ")
    total := 0
    for _, num := range nums {
        total += num
    }
    fmt.Println(total)
}

func main() {

    // 가변인자함수는 독립적인 파라미터들과 함께 호출될수 있다. 
    sum(1, 2)
    sum(1, 2, 3)

    //
    //슬라이스안에 이미 여러개의 파라미터를 가진 경우 그것들을 다음과 같이 
    //메서드(슬라이스)같이 하여 가변인자함수에 적용할 수 있다.  
    nums := []int{1, 2, 3, 4}
    sum(nums...)
}
