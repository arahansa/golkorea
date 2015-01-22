/*
    author twitter : @mmcgrana
	https://gobyexample.com/worker-pools
    Go by Example: worker-pools


	이 예제에서 우리는 어떻게 go루틴과 채널을 이용하여 worker pool을 구현할지 살펴보자.
*/

package main

import "fmt"
import "time"

//여기 worker 가 있는데 우리는 동시성 인스턴스 몇가지를 실행시킬 것이다.
//이러한 worker들은 jobs채널로부터 work를 받아서 그에 상응하는 results를 보낼 것이다.
//우리는 처리하는데 시간이 걸리는(expensive)한 job이란 시뮬레이션을 하기 위해 job한개당 일초간 Sleep 할 것이다.
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "processing job", j)
		time.Sleep(time.Second)
		results <- j * 2
	}
}

func main() {

	// workers의 pool을 사용하기 위해 우리는 그들에게 work를 보내고 그들의 결과를 모을 것이다.
	//우리는 이를 위해 2개의 채널을 만들 것이다.
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// 이것은 3 개의 workers를 실행시킬 것이고 처음엔 아무 jobs도 없기 때문에 블락될것이다.
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	//여기 우리는 9개의 jobs를 보내고 이것이 우리가 가진 모든 work란 것을 알리기 위해 저 채널을 닫을 것이다.
	for j := 1; j <= 9; j++ {
		jobs <- j
	}
	close(jobs)
	//마침내 우리는 work의 모든 결과들을 모을 것이다.
	for a := 1; a <= 9; a++ {
		<-results
	}
}
