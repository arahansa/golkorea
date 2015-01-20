/*
원글 링크 : https://code.google.com/p/golang-korea/wiki/EffectiveGo
명명된 파라미터 결과값(Named result parameters)
Go 함수에서는 리턴값 또는 parameter 결과값들에 대해 명명하는 것이 가능하며
입력되는 parameter들 처럼 일반적인 변수로 사용하는 것 또한 가능합니다.

리턴값을 명명하게 되면 그 함수가 시작될 때 리턴값은 리턴값의 타입에 대해
zero 값으로 초기화 됩니다. 만약 함수가 argument 없이 리턴문을 실행하게 되면
현재 parameter들의 결과값을 리턴값으로써 사용하게 됩니다.

결과값에 이름을 짓는 것은 필수 사항이 아니지만, 리턴값에 명명하는 것은 코드를
좀 더 간단하고 명료하게 할 수 있습니다.
: 만약 위의 예제에서 nextInt의 결과값을 이름을 붙여주게 되면
어떤 int값이 리턴되어 졌는지 확실히 알게 될 것입니다.

func nextInt(b []byte, pos int) (value, nextPos int) {
위의 예에서 보듯이 이름붙여진 결과값들이 초기화되었고 간소화되었기 때문에,
단순할 뿐만아니라 명확해질 수 있습니다.
여기 자주 사용되는 io.ReadFull?의 예가 있습니다.
*/
func ReadFull(r Reader, buf []byte) (n int, err os.Error) {
	for len(buf) > 0 && err == nil {
		var nr int
		nr, err = r.Read(buf)
		n += nr
		buf = buf[nr:len(buf)]
	}
	return
}