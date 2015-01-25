/*
원본 주소 : 
https://code.google.com/p/golang-korea/wiki/EffectiveGo#메소드(Methods)



메소드(Methods)
포인터 vs 값(Pointers vs. Values)
메소드는 pointer나 interface가 아니라면 어떤 이름의 타입으로든지 정의되어 질 수 있습니다. 
수신자(receiver)는 구조체(struct)여서는 안됩니다. 

위에 slice에 대해 토론한 부분에서, append 함수에 대해서 언급했었습니다. 
대신 우리는 슬라이스에서 그것을 메소드로 정의할 수도 있습니다. 
이렇게 하기 위해서, 우리는 먼저 메소드와 묶을 수 있도록 타입을 정의해야 합니다. 
그리고, 그 타입의 값인 메소드에 대한 수신자(receiver)를 만들어야 합니다.
*/
type ByteSlice []byte

func (slice ByteSlice) Append(data []byte) []byte {
    // Body는 위에서 했던것과 똑같습니다.
}

//이 예제는 여전히 업데이트된 slice를 반환하기 위한 방법이 필요합니다. 
//우리는 수신자(receiver)로서 Byteslice의 포인터를 받아들이도록 메소드를 다시 정의함으로써 
//어색한 부분을 고칠 수 있습니다. 그렇게하면 메소드는 호출자(caller)의 slice를 덮어쓸 수 있습니다.
func (p *ByteSlice) Append(data []byte) {
    slice := *p
    // Body는 위와 같지만 return이 없습니다.
    *p = slice
}
//사실 우리는 이것보다 더 잘할 수 있습니다. 
//만약 우리가 우리의 함수를 수정한다면 아래 예제처럼 그것은 표준적인 메소드처럼 보일 것 입니다.
func (p *ByteSlice) Write(data []byte) (n int, err os.Error) {
    slice := *p
    // 위의 예제와 같음
    *p = slice
    return len(data), nil
}
//그렇게되면, *ByteSlice 타입은 표준 인터페이스인 io.Writer를 만족시키며 편리합니다. 
//예를들면, 우리는 아래 예처럼 한번에 출력할 수 있습니다.

    var b ByteSlice
    fmt.Fprintf(&b, "This hour has %d days\n", 7)
//우리는 ByteSlice?의 주소값을 넘겼습니다. 왜냐하면 *ByteSlice만이 io.Writer에 적합하기 때문입니다. 
//수신자(receiver)에 대한 포인터와 값에 대한 규칙은 값 메소드는(value method) 포인터와 값을 불러올 수 있지만 
//포인터 메소드는(pointer method) 오로지 포인터만을 불러올 수 있다는 것 입니다. 
//이것은 포인터 메소드는 수신자(receiver)를 수정할 수 있기 때문입니다. 
//값의 복사본에서 그것들을 불러오는 경우는 이런 수정들이 적용되지 않을 것 입니다.
//byte slice에서 사용되어진 아이디어는 bytes.Buffer에 구현되어 있습니다.