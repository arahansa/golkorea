/*
원글 링크 : https://code.google.com/p/golang-korea/wiki/EffectiveGo
다중 리턴 값(Multiple return values)
Go언어에서 특이한 점 중에 하나는 함수와 메소드가 여러개의 값을 리턴할 수 있다는 것입니다. 
이런 형태는 C 프로그램에서 몇몇 어색한 구문들을 개선하는데 사용될 수 있습니다.
: in-band 에러 리턴(EOF를 위한 -1과 같은) 과 argument의 수정

C언어에서 쓰기 에러는 임시 주소에서 에러 코드와 함께 음의 숫자로 표현됩니다.
Go언어에서는 숫자와 에러를 리턴할 수 있습니다
:"몇몇 byte를 쓸 수 있지만 그렇게하면 device 전체 공간을 체우게될 수도 있기 때문에 전부 다 쓸 수는 없습니다." 
os 패키지에서 *File.Write의 signature는 다음과 같습니다.
*/

func (file *File) Write(b []byte) (n int, err Error)
/*
그리고 문서에서 말하는 것 처럼, 이 함수는 쓰여진 byte의 수와 n != len(b)인 경우 
non-nil 에러를 리턴할 것 입니다. 이것은 일반적인 형태입니다; 

더 많은 예제를 확인하고 싶으면 error handling section을 참고하십시오.

비슷한 접근법으로 parameter들을 참조하기 위해 리턴 값으로 부터 포인터를 사용하는 것을 피할 수 있습니다. 
다음은 byte array의 한 위치로 부터 숫자를 리턴하고 다음 위치로 이동하기 위한 간단한 함수 예제입니다.
*/
func nextInt(b []byte, i int) (int, int) {
    for ; i < len(b) && !isDigit(b[i]); i++ {
    }
    x := 0
    for ; i < len(b) && isDigit(b[i]); i++ {
        x = x*10 + int(b[i])-'0'
    }
    return x, i
}
/*
입력된 배열에서 숫자를 검색하기 위해 다음과 같이 할 수 있습니다.
*/
    for i := 0; i < len(a); {
        x, i = nextInt(a, i)
        fmt.Println(x)
    }
	