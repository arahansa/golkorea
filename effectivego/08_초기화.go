/*
비록, C나 C++의 초기화 과정과 표면적으로 아주 다르게 보이는 부분은 없지만,
Go 언어의 초기화 과정은 보다 강력합니다.
초기화 과정에서 복잡한 구조체가 만들어지고,
초기화된 Object와 여타 package간의 배치가 보다 명확하게 다루어집니다.

상수(Constants)
Go 언어에서 상수는 말그대로 상수입니다.
상수는 컴파일 시점에 생성되며, - 심지어 함수의 지역 상수라도 - 그 값은 숫자, 문자열, boolean 만 가능합니다.
컴파일 시점에 생성되는 제약때문에, 상수를 정의하는 표현은 컴파일러가 이해할 수 있는 상수 형태이어야 합니다.
예를 들어 1<<3 은 상수 형태이고, math.Sin(math.Pi/4)는 math.Sin 함수가 런타임에 호출되어야 하므로 상수 형태라고 할 수 없습니다.

Go 언어에서, 열거된 상수는 iota 열거자(iota enumerator)를 사용해서 생성합니다.
iota는 expression의 한 부분이 될 수 있고, expression은 무조건적인 반복이 될 수 있기 때문에,
그것은 복잡한 값의 집합이 되기 쉽습니다.
*/
type ByteSize float64

const (
	_           = iota // ignore first value by assigning to blank identifier
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

/*
String과 같은 메소드를 추가하는 기능은 일반적인 타입의 한 부분일지라도,
출력에 대해 그런 값들이 자동으로 형식화되도록 하는 것을 가능하게 합니다.
*/
func (b ByteSize) String() string {
	switch {
	case b >= YB:
		return fmt.Sprintf("%.2fYB", float64(b/YB))
	case b >= ZB:
		return fmt.Sprintf("%.2fZB", float64(b/ZB))
	case b >= EB:
		return fmt.Sprintf("%.2fEB", float64(b/EB))
	case b >= PB:
		return fmt.Sprintf("%.2fPB", float64(b/PB))
	case b >= TB:
		return fmt.Sprintf("%.2fTB", float64(b/TB))
	case b >= GB:
		return fmt.Sprintf("%.2fGB", float64(b/GB))
	case b >= MB:
		return fmt.Sprintf("%.2fMB", float64(b/MB))
	case b >= KB:
		return fmt.Sprintf("%.2fKB", float64(b/KB))
	}
	return fmt.Sprintf("%.2fB", float64(b))
}

/*
( float64 변환은 Sprintf가 ByteSize?에 대해 String 메소드를 통해 반복되는 것을 막아줍니다.)
YB 표현식(expression)은 1.00YB로 출력됩니다. 반면 ByteSize?(1e13)은 9.09TB로 출력됩니다.

변수(Variables)

변수는 상수처럼 초기화될 수 있지만 초기화는 런타임(run time)에 계산되는 일반적인 표현식(expression)이 될 수도 있습니다.
*/
var (
	HOME   = os.Getenv("HOME")
	USER   = os.Getenv("USER")
	GOROOT = os.Getenv("GOROOT")
)

/*
초기화 함수(The init function)
마침내, 각 소스파일은 어떤 명령문이던 간에 필요하다면 셋업하기 위해 자신만의 init() 함수를 정의할 수 있습니다.
유일한 제한은 비록, 초기화하는 동안 go 루틴이 시작될 수 있다고 해도, 초기화가 끝날때까지는 실행되지 않을 것이라는 점입니다.
초기화는 항상 단일 스레드로 실행됩니다. 그리고 마지막으로, 팩케지 안에 모든 변수 선언이 초기화된 뒤에 init()이 호출되어 집니다.

선언으로서 표현될 수 없는 초기화에 비해, init() 함수의 일반적인 사용은 실제 실행이 시작되기 전에
검증이나 프로그램 명령문의 정확한 수정을 위해 사용될 수 있습니다.
*/
func init() {
	if USER == "" {
		log.Fatal("$USER not set")
	}
	if HOME == "" {
		HOME = "/usr/" + USER
	}
	if GOROOT == "" {
		GOROOT = HOME + "/go"
	}
	// GOROOT may be overridden by --goroot flag on command line.
	flag.StringVar(&GOROOT, "goroot", GOROOT, "Go root directory")
}