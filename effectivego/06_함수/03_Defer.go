/*Defer
원글 링크 : https://code.google.com/p/golang-korea/wiki/EffectiveGo
Go 언어에서 defer 문은 함수가 지연된 리턴값을 실행하기 전에 즉시 실행될 수 있도록 
하나의 function call을 예약합니다.(the deferred function) 
일반적이지는 않지만 어떤 경로의 함수가 값을 리턴하는지에 관계없이 자원을 해제해야만하는 상황을 
다루기 위해서는 효과적인 방법입니다.전형적인 예로 mutex를 해제하거나 file을 닫는 경우 입니다.

// Contents returns the file's contents as a string.
*/
func Contents(filename string) (string, os.Error) {
    f, err := os.Open(filename, os.O_RDONLY, 0)
    if err != nil {
        return "", err
    }
    defer f.Close()  // f.Close()는 현재의 함수가 종료되면 실행되어 질 것입니다.

    var result []byte
    buf := make([]byte, 100)
    for {
        n, err := f.Read(buf[0:])
        result = append(result, buf[0:n]...) // append와 관련해서는 추후에 다루게 될 것입니다.
        if err != nil {
            if err == os.EOF {
                break
            }
            return "", err  // 여기서 리턴한다면 f는 닫히게 될 것입니다.
        }
    }
    return string(result), nil // 여기서 리턴한다면 f는 닫히게 될 것입니다.
}
/*
Close와 같은 함수에서 지연 호출은 2가지 장점이 있습니다. 
첫번째로, 파일을 닫는 것을 잊지 않도록 해 줄 것입니다.
(만약 함수에서 뒤에 파일 닫는 부분을 작성한다면 잊기 쉽습니다.) 

두번째로, 이것은 close가 open 근처에 놓이게 된다는 것을 의미하며 함수의 뒷쪽에 놓이는 것보다 
훨씬 더 분명해 보입니다. 지연 함수(the deferred function)에서 
인자값들은(만약 함수가 메소드라면 수신자(receiver)를 포함해서) 호출이 실행될 때가 아니라 
defer문이 실행되어 질 때 평가되어 집니다. 

반면, 함수 실행시에 변수값들의 변화에 대한 문제를 피한다는 측면에서, 
이것은 하나의 지연된 호출이 여러개의 함수 실행을 지연시킬 수 있다는 의미를 가집니다. 
여기 아주 쉬운 예가 있습니다.
*/
for i := 0; i < 5; i++ {
    defer fmt.Printf("%d ", i)
}

/*
지연된 함수들은 LIFO의 순서로 실행되어 집니다. 
그러므로 이 코드는 함수가 종료되었을 때 4 3 2 1 0의 순서로 프린트되어질 것입니다.

좀 더 괜찮은 예로 프로그램을 통해 함수 실행을 추적해보는 쉬운 방법이 있습니다. 
우리는 다음과 같이 단순히 경로를 추적하는 몇개의 코드를 작성할 수 있습니다.
*/
func trace(s string)   { fmt.Println("entering:", s) }
func untrace(s string) { fmt.Println("leaving:", s) }

// Use them like this:
func a() {
    trace("a")
    defer untrace("a")
    // do something....
}
/*
우리는 지연된 함수에서 인자값들이 defer문이 실행되어질 때 평가되어진다는 사실을 이용해서 
더 좋은 결과를 얻을 수 있습니다. 경로 추적은 루틴에 인자값을 정함으로서 할 수 있습니다.
*/
func trace(s string) string {
    fmt.Println("entering:", s)
    return s
}

func un(s string) {
    fmt.Println("leaving:", s)
}

func a() {
    defer un(trace("a"))
    fmt.Println("in a")
}

func b() {
    defer un(trace("b"))
    fmt.Println("in b")
    a()
}

func main() {
    b()
}
/*
prints

entering: b
in b
entering: a
in a
leaving: a
leaving: b
프로그래머들이 다른 언어에서는 차단되어져있는(block-level) 자원의 관리에 익숙하기 때문에, 
defer문은 이상하게 보일지도 모릅니다.하지만 defer문의 가장 흥미롭고 강력한 적용은 block기반이 아니라
function 기반으로부터 온다는 것 입니다. 

혼란과 복구(panic and recover) 부분에서 우리는 defer문의 또 다른 예를 보게 될 것 입니다.

*/