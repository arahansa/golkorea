/*
    author twitter : @mmcgrana
    https://gobyexample.com/defer

	Go by Example: defer

	Defer가 프로그램의 실행 중 함수호출이수행된 후에 cleanup 되었단 것을 보장하기 위해 사용된다.
	defer 는 종종 다른 언어에서 finally나 ensure 같은 것처럼 사용된다.
*/

package main

import "fmt"
import "os"

// 우리가 새로 파일을 생성하고 작성을 한뒤 우리의 작업을 마치고 닫는다고 가정하자.
// 여기 우리가 어떻게 defer와 함께 그 일을 하는 지 나와있다.
func main() {

	//  파일객체를 createFile로 받은 즉시, 우리는  그 파일을  닫는 것을 'closeFile'와 함께 연기(defer)할 것이다.
	// writeFile를 마친 후, 이것을 둘러싼 함수(main함수)가 끝나게 되면서 defer는 실행될 것이다.
	f := createFile("/tmp/defer.txt")

	//f := createFile("c:/gocode/defer.txt") 윈도우의 경우 다음과 같이 해보자(역주)
	defer closeFile(f)
	writeFile(f)
}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")

}

func closeFile(f *os.File) {
	fmt.Println("closing")
	f.Close()
}

// 프로그램을 실행하면 쓰기작업후에 file이 닫혀졌는지 확인한다.
