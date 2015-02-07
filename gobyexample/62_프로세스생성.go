/*
    author twitter : @mmcgrana
   	https://gobyexample.com/spawning-processes
	Go by Example:spawning-processes


	때때로 우리의 Go 프로그램들은 Go 가 아닌 프로세스같은 것들을 만들기 원한다.
	예를 들자면 이 사이트의 신택스 Go 프로그램에서 pygmentize를 spawning함으로써 구현되었다.
	Go로부터 spawning 프로세스들의 예제를 한번 살펴보자.
*/

package main

import "fmt"
import "io/ioutil"
import "os/exec"

func main() {

	// 우리는 어떠한 아규먼트나 입력값을 받지 않는 간단한 커맨드로 시작해보고자 한다.
	//그리고  어떤 것을 stdout에 출력시킬 것이다.
	// 'exe.Command' 헬퍼는 객체를 생성하여 이러한 외부 처리를 표현한다.
	dateCmd := exec.Command("date")

	// .Output 은 다른 다른 헬퍼로써, 보통 실행중인 커맨드를 다루는 경우나, 그것이 마칠때까지 기다리거나
	// 그것의 결과를 모은다. 에러가 없다면  dateOut은 날짜 정보와 함께 바이트를 가지고 있을 것이다.
	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> date")
	fmt.Println(string(dateOut))

	// 다음 우리는 좀더 연관된 케이스를 살펴 볼 것인데 이 케이스는
	// 외부 프로세스의 stdin으로 데이터를 보내고
	// 결과를 그것의 stdout으로 받을 것이다.
	// Next we'll look at a slightly more involved case
	// where we pipe data to the external process on its
	// `stdin` and collect the results from its `stdout`.
	grepCmd := exec.Command("grep", "hello")

	// 여기 우리는 명시적으로 input/output 파이프들을 잡을 것이다.
	// 프로세스를 시작하면서 몇몇 입력을 여기에 할 것이고
	// 결과 출력을 읽을 것이며, 최종적으로 프로세스가 종료되길 기다릴 것이다.
	// Here we explicitly grab input/output pipes, start
	// the process, write some input to it, read the
	// resulting output, and finally wait for the process
	// to exit.
	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	grepCmd.Start()
	grepIn.Write([]byte("hello grep\ngoodbye grep"))
	grepIn.Close()
	grepBytes, _ := ioutil.ReadAll(grepOut)
	grepCmd.Wait()

	// 우리는 위의 예제에서 에러체크를 생략했다. 그러나 우리는 흔한 'if err != nil '패턴을 모든 이것들에 써도 된다.
	// 우리는 또한 'StdoutPipe'결과귿릉ㄹ 모을 것이나, 당신은 또한 정확히 같은 방법으로 'StderrPipe'를 받아도 된다.
	// We ommited error checks in the above example, but
	// you could use the usual `if err != nil` pattern for
	// all of them. We also only collect the `StdoutPipe`
	// results, but you could collect the `StderrPipe` in
	// exactly the same way.
	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))

	// 커맨드를 실행할 때는 우리가 기술된 커맨드와 아규먼트들을  제공할 필요가 있다는 것을 기억하자
	// 반면에 하나의 커맨드-라인 문자열을 전달해도 된다.
	// 만약 당신이 완전한 커맨드를 문자열로 실행시키고 싶으면 당신은 'bash'의 '-c'옵션을 사용할 수 있다
	// Note that when spawning commands we need to
	// provide an explicitly delineated command and
	// argument array, vs. being able to just pass in one
	// command-line string. If you want to spawn a full
	// command with a string, you can use `bash`'s `-c`
	// option:
	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut))
}
