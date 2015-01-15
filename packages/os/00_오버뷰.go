/*
os 패키지는 운영체제 기능에 대하여 플랫폼-독립적인 인터페이스를 제공합니다.
에러핸들링은 Go같이 되었지만. 디자인은 유닉스와 같이 되었습니다.

실패된 호출은 에러 넘버들을 리턴하기보다는, 타입에러 값을 리턴합니다.
종종, 많은 정보들이 에러 내에서 사용가능합니다.

예를 들자면, Open 이나 Stat같은 파일이름으로 하는 호출이 실패하였을 경우,
에러가 출력될 때, 실패한 파일이름을 포함하고 에러는 더 많은 정보를 위해 unpack 된 *PathError 타입이 될 것입니다.

os 인터페이스는 모든 운영체제에 대하여 유연하게 설계되었습니다.(일부러 사전적단어를 바꾼 번역함 intended to be uniform)
Features not generally available appear in the system-specific package syscall.
(이건 뭔소리인지 감이 잘 안오네요^^)

여기 파일을 열고 읽는 예제가 있습니다.
*/
file, err := os.Open("file.go") // For read access.
if err != nil {
	log.Fatal(err)
}
/*
열기에 실패하면, 에러문자열은 설명적으로 다음과 같이 나타날 것이다.*/

open file.go: no such file or directory

/*
파일의 데이터는 byte의 슬라이스로 읽혀질 것입니다.. 
읽고쓰는 작업은 slice의 길이로부터 그 파일의 byte counts를 가질 것입니다.
*/
data := make([]byte, 100)
count, err := file.Read(data)
if err != nil {
	log.Fatal(err)
}
fmt.Printf("read %d bytes: %q\n", count, data[:count])


