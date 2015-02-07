/*
    author twitter : @mmcgrana
   	https://gobyexample.com/execing-processes
	Go by Example: execing-processes

	이전의 예제에서 우리는 외부 프로세스들을 생성하는 법을 살펴보았는데,
	우리는 실행중인 Go 프로그램에서 접근가능한 외부프로세스가 필요할 때 이런 것을 할 수 있다.
	때때로 우리는 현재의 Go프로세스를 Go프로세스가 아닌 다른 것으로 대체하고 싶을 때가 있다.
	이러한 것을 하기 위해, 우리는 Go의 exec함수를 사용할 것이다.
*/

// In the previous example we looked at
// [spawning external processes](spawning-processes). We
// do this when we need an external process accessible to
// a running Go process. Sometimes we just want to
// completely replace the current Go process with another
// (perhaps non-Go) one. To do this we'll use Go's
// implementation of the classic
// <a href="http://en.wikipedia.org/wiki/Exec_(operating_system)"><code>exec</code></a>
// function.

package main

import "syscall"
import "os"
import "os/exec"

func main() {

	// 우리의 예제에서, 우리는 ls를 실행할 것이다. Go는 우리가 실행하고자 하는 binary에의 절대경로를 필요로 한다.
	// 그래서 우리는 exec.LookPath를 사용하여서 그것을 찾아볼 것이다.
	// For our example we'll exec `ls`. Go requires an
	// absolute path to the binary we want to execute, so
	// we'll use `exec.LookPath` to find it (probably
	// `/bin/ls`).
	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}

	// 'Exec'는 나란히 한 줄로 놓여진 슬라이스 형식의 아규먼트를 필요로 한다.
	// 우리는 ls에게 몇가지 흔한 아규먼트들을 줄 것이다.
	// 첫번째 아규먼트는 프로그램 이름이 되어야 한다는 것을 기억하자.
	args := []string{"ls", "-a", "-l", "-h"}

	// 'Exec'는 또한 환경변수 모음이 필요하다
	// 여기 우리는 우리의 환경변수를 줄 수 있다.
	// `Exec` also needs a set of [environment variables](environment-variables)
	// to use. Here we just provide our current
	// environment.
	env := os.Environ()

	// 여기 실제 os.Exec호출이 있다. 만약 이 호출이 성공적이라면
	// 우리의 프로세스의 실행은 여기서 끝나고 /bin/ls -a -l -h 프로세스가 대신할 것이다.
	// 만약 에러가 있다면, 우리는 리턴값을 얻을 것이다.
	// Here's the actual `os.Exec` call. If this call is
	// successful, the execution of our process will end
	// here and be replaced by the `/bin/ls -a -l -h`
	// process. If there is an error we'll get a return
	// value.
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}
