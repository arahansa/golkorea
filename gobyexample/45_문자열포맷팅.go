/*
    author twitter : @mmcgrana
    https://gobyexample.com/string-formatting
	Go by Example: string-formatting

	Go는 printf전통속에서 문자열 형식을 훌륭하게 지원한다.
	여기 평범한 문자열형식처리의 예제가 있다.
*/

package main

import "fmt"
import "os"

type point struct {
	x, y int
}

func main() {

	//Go는 몇몇 printing "동사"를 제공하는데 이것은 일반적 Go의 값들을 형식화하기 위해서다.
	//예를 들자면 이것은 point구조체의 인스턴스를 출력한다.
	p := point{1, 2}
	fmt.Printf("%v\n", p)

	//만약 값이 구조체라면 %+v 변종은 구조체의 name또한 포함한다.
	fmt.Printf("%+v\n", p)

	// %#v 변종은 값의 Go문법표현을 출력한다. 예를 들어서 그 값을 만들어내는 소스코드 snippet 같은 것들말이다.
	fmt.Printf("%#v\n", p)

	// 값의 타입을 출력하기 위해 %T를 사용했다.
	fmt.Printf("%T\n", p)

	// 참거짓형식화는 간단하다.
	fmt.Printf("%t\n", true)

	// 정수를 형식화하는 데는 많은 옵션이 있는데 %d가 표준이고 10진수다.
	fmt.Printf("%d\n", 123)

	// 바이러니 표현을 출력한다.
	fmt.Printf("%b\n", 14)

	//주어진 값에 맞는 문자를 출력한다.
	fmt.Printf("%c\n", 33)
	// %x는 hex인코딩을 제공한다.
	fmt.Printf("%x\n", 456)

	// Float를 위한 몇몇 옵션또한 있다.
	//기본적인 Decimal 포맷은 %f사용이다.
	fmt.Printf("%f\n", 78.9)

	// %e나 %E 는(약간 다른 버젼이다) 과학적 표시에서의 형식이다.
	fmt.Printf("%e\n", 123400000.0)
	fmt.Printf("%E\n", 123400000.0)

	// 기본 문자열출력은 %s를 사용
	fmt.Printf("%s\n", "\"string\"")

	// Go 소스 에서 두개 더블 인용문자열을 하려면%q를 사용하라.
	fmt.Printf("%q\n", "\"string인용구 테스트\"")

	// 우리가 전에 본 정수처럼, %x는 입력바이트당 두개의 출력문자열들을 가지고서 16진수문자열을 렌더링한다.
	fmt.Printf("%x\n", "hex this")

	// 포인터 표현(representation)을 출력하기 위해 여기선 %p를 썼다.
	fmt.Printf("%p\n", &p)

	// 숫자를 형식화할때, 종종 결과수치의 가로길이를 종종 조정하고 싶을 지도 모른다.
	// 가로길이를 정하기 위해, %뒤에 숫자를 써준다.
	// 디폴트로 결과를 오른쪽정렬될 것이며 왼쪽에 공백(왼쪽)과 함께 들어갈 것이다.
	fmt.Printf("|%6d|%6d|\n", 12, 345)

	// 당신은 또한 출력된 실수의 가로길이를 정할 수 있는데, 또한 실수 소수점 뒷자리 부분길이도 정하고 싶을 것이다.
	// 그럴때는 길이.소수점뒷자리수길이 로 해주면 된다.
	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)

	// 왼쪽 정렬을 하기 위해 - 플래그를 이용하자
	fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)

	// 문자열 포맷팅 할때 가로길이를 정하고 싶을 지도 모른다. 특히 테이블비슷한 출력을 하고 싶을 때 말이다.
	// 다음은 오른쪽 정렬되는 문자열들이다.
	fmt.Printf("|%6s|%6s|\n", "foo", "b")

	//왼쪽정렬을 할때는 - 플래그를 써준다.
	fmt.Printf("|%-6s|%-6s|\n", "foo", "b")

	//지금까지 우리는 'Printf'를 살펴보았는데, 이것은 형식화된 문자열을 'os.Stdout'에 출력한다
	// Sprintf는 문자열을 형식화하며, 출력을 하지 않고 문자열을 리턴한다.
	s := fmt.Sprintf("a %s", "string")
	fmt.Println(s)

	//또한 Fprintf를 사용하여 형식화+출력을 'os.Stdout'이 아닌 'io.Writers'에 할 수 있다.
	fmt.Fprintf(os.Stderr, "an %s\n", "error")
}
