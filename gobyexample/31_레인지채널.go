/*
    author twitter : @mmcgrana
	https://gobyexample.com/range-over-channels
    Go by Example: range-over-channels

	지난번의 range예제에서 우리는 for와 range가 기초적인 자료구조에서 어떻게 반복하는지 보았다.
	우리는 또한 이 문법을 채널을 통해 들어오는 값에서 반복할 수 있다.
*/

package main

import "fmt"

func main() {

	// 우리는 queue 채널에서 2개의 값 이상을 반복할 것이다.
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	//이 range는 각각의 요소들 위에서 반복한다. 요소들이 queue채널에서 받아지기 때문이다.
	//우리가 위에서 채널을 닫았기 때문에 반복은 2개의 요소를 받고 난 뒤 끝난다.
	// 만약 우리가 채널을 닫지 않으면, 우리는 loop 에서 3번째로 받을 때 블락될 것이다.

	for elem := range queue {
		fmt.Println(elem)
	}
}
