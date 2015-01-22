/*

원글 링크 : https://code.google.com/p/golang-korea/wiki/EffectiveGo


출력 (Printing)
Go언어에서 일반적인 출력은 C언어에서 printf와 비슷한 형태를 사용합니다. 
하지만 더 많은 옵션을 제공하며 일반화되어 있습니다. 
출력 관련 함수들은 fmt 펙케지 안에 있으며 대문자 이름으로 되어 있습니다: 
fmt.Printf, fmt, Fprintf, fmt.Sprintf 등등. string 관련 함수들(Sprintf 등)은 
제공된 버퍼를 체우기 보다는 string을 반환합니다.

형식적인 string을 제공할 필요가 없습니다. 
각 Printf, Fprintf, Sprintf 함수에는 또 다른 쌍의 함수가 있습니다. 
예를 들면, Print와 Println 입니다. 이러한 함수들은 형식화된 string을 취하지 않지만 
대신, 각각의 인자로 기본 형태의 string을 만듭니다. 
Println 함수는 인자와 새로운 라인을 추가하는 사이에 빈공간을 추가합니다. 
반면, Print 함수의 경우는 오직 양쪽 피연산자가 둘다 string이 아닌 경우만 blank를 추가 합니다. 
여기 예제에서 각 라인은 같은 결과를 출력합니다.
*/

fmt.Printf("Hello %d\n", 23)
fmt.Fprint(os.Stdout, "Hello ", 23, "\n")
fmt.Println("Hello", 23)
fmt.Println(fmt.Sprint("Hello ", 23))
/*tutorial에서 말했던 것 처럼, fmt.Fprint와 그와 비슷한 함수들은 
io.Writer 인터페이스를 구현하는 어떤 오브젝트의 첫번째 인자처럼 받아들입니다. 
os.Stdout과 os.Stderr 변수들과 비슷한 경우 입니다.

여기 있는 예제는 C언어와 달라지기 시작하는 것을 보여줍니다. 
첫번째로, %d와 같은 숫자 형식은 부호나 크기를 위한 flag를 취하지 않습니다. 
대신 출력 루틴은 이런 속성들을 결정하기 위해 인자 타입을 사용합니다.
*/

var x uint64 = 1<<64 - 1
fmt.Printf("%d %x; %d %x\n", x, x, int64(x), int64(x))
prints

18446744073709551615 ffffffffffffffff; -1 -1
/*
만약 10진수 정수와 같은 기본 변환을 원한다면, 포괄적인 형태의 %v를 사용할 수 있습니다.
(값을 위해서) 결과는 Print와 Println이 만들어 내는 것과 정확히 같습니다. 
더욱이 그런 형식으로 어떤 값이든 출력할 수 있습니다.(심지어 배열, 구조체,map 까지도) 
여기 앞의 예제에 정의되어 있는 time zone map을 사용하는 print문 예제가 있습니다.

fmt.Printf("%v\n", timeZone)  // or just fmt.Println(timeZone)

출력:
map[CST:-21600 PST:-28800 EST:-18000 UTC:0 MST:-25200]
물론, map에서 key는 순서없이 출력될 것 입니다. 구조체를 출력할 때, 변형 포맷인 %+v는 
구조체 필드의 이름을 표시할 것입니다. 그리고 어떤 값에 대해서든 %#v 형식은 전체 Go 구문을 출력할 것 입니다.
*/
type T struct {
    a int
    b float
    c string
}
t := &T{ 7, -2.35, "abc\tdef" }
fmt.Printf("%v\n", t)
fmt.Printf("%+v\n", t)
fmt.Printf("%#v\n", t)
fmt.Printf("%#v\n", timeZone)
prints

&{7 -2.35 abc   def}
&{a:7 b:-2.35 c:abc     def}
&main.T{a:7, b:-2.35, c:"abc\tdef"}
map[string] int{"CST":-21600, "PST":-28800, "EST":-18000, "UTC":0, "MST":-25200}
/*
(&기호에 주목하십시오.) 인용부호로("") 둘러쌓인 문자열 형식은 string 또는 byte 배열(byte) 형태의 
값을 적용할 때 %q를 통해서도 가능 합니다. %#q 형식은 만약 가능하다면 인용부호("") 대신 backquote(`)를 
사용할 것 입니다. 또한 %x 옵션은 integer뿐만 아니라 문자열과 byte 배열에서도 동작 합니다. 

그것은 긴 16진 문자열을 생성하며, %x에서 space를 사용하는 것(% x)은 byte 사이에 space를 넣게 됩니다.
또다른 유용한 형태는 값의 타입을 출력하는 %T 옵션 입니다.
*/
fmt.Printf("%T\n", timeZone)
prints

map[string] int
/*
만약 custom 타입을 기본 형태로 제어하기를 원한다면, 모든 요구사항은 
String() string 메소드에 그 타입을 정의해야 합니다. T 타입은 다음 예제와 같이 보여질 것 입니다.
*/
func (t *T) String() string {
    return fmt.Sprintf("%d/%g/%q", t.a, t.b, t.c)
}
fmt.Printf("%v\n", t)
/*
to print in the format

7/-2.35/"abc\tdef"
String() 메소드는 Sprintf를 호출 하는 것이 가능합니다. 왜냐하면 print 루틴은 재귀적으로 사용될 수 있기 때문입니다.

이제 한단게 더 나갈 수 있습니다. 그리고 print 루틴의 인자들을 바로 다른 루틴에 넘겨줄 수 있습니다. 
Printf 함수의 특징은 포맷 이후에 나타날 수도 있는 임의의 개수의 파라미터(또는 임의의 형태의 파라미터)들을 
명시하기 위해 마지막 인자로 ...interface{} 타입을 사용한다는 것 입니다.

*/
func Printf(format string, v ...interface{}) (n int, errno os.Error) {
/*
Printf 함수 내부에서, v는 interface{} 타입의 변수처럼 동작합니다. 하지만 만약 또 다른 가변인자 함수로 넘겨진다면, 
그것은 일반적인 인자 리스트 처럼 동작할 것 입니다. 여기 log.Println을 구현한 예제가 있습니다.(위에서 사용한) 
이 함수는 실제 설정을 위해 인자값들을 바로 fmt.Sprintln 함수로 넘깁니다.


// Println prints to the standard logger in the manner of fmt.Println.
*/
func Println(v ...interface{}) {
    std.Output(2, fmt.Sprintln(v...))  // Output takes parameters (int, string)
}
/*
인자 리스트로서 v를 취급하도록 컴파일러에 알려주기 위해, Sprintln에 대한 내부 호출에서 v 뒤에 ...을 사용 합니다. 
그렇지 않다면, 단순히 v를 싱글 slice 인자로 넘겨줄 것 입니다.

출력 관련해서는 이 문서에서 다루고 있는 것 이상의 내용이 있습니다. 
fmt 팩키지에 대한 더 자세한 내용을 알고 싶다면 godoc 문서를 보십시오.

어쨌거나 ... 파라미터는 특정 타입이 될 수 있습니다. 예를 들면 integer 
리스트에서 최소값을 고르는 min 함수를 위해 ...int를 사용할 수 있습니다.
*/
func Min(a ...int) int {
    min := int(^uint(0) >> 1)  // largest int
    for _, i := range a {
        if i < min {
            min = i
        }
    }
    return min
}