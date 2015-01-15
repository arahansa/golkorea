/*

https://code.google.com/p/golang-korea/wiki/EffectiveGo#명칭(Names)

Go 언어는 getter와 setter를 자동으로 제공하지 않습니다.
당신 스스로 getter와 setter를 만드는 것은 잘못된 것이 아니며,
그렇게 하는 것은 흔한 일입니다.

당신이 만든 getter와 setter에는 문법적인 오류가 없어야 하며
getter의 이름에 Get을 넣을 필요가 없습니다.

예로, 만약 당신이 owner(소문자, export되지 않았음)라고 불리는 field를 하나 가지고 있다면
getter method는 GetOwner?가 아니라 Owner(대문자, export됨)로 호출되어야만 합니다.

export하기 위해 대문자 이름을 사용하는 것은 method로 부터 그 field를 구분하기 위한 방법을
제공합니다. 만약 필요하다면, setter 함수는 SetOwner?로 불릴 것입니다.

두 이름 모두 읽기에 좋습니다.(getter는 Get이라는 단어를 사용하지 않지만,
setter는 Set이라는 단어를 사용함)

*/
func main() {
	owner := obj.Owner()
	if owner != user {
		obj.SetOwner(user)
	}
}
