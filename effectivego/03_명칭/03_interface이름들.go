/*

Interface 이름들
규정에 의하면 하나의 method interface들은 method name에 추가로 -er 접미사를 붙여서
이름을 만듭니다.(예: Reader, Writer, Formatter 등등)

그런류에 많은 이름들이 있고 이런 방법은 그것들을 관리하는데 효율적입니다.
Read, Write, Close, Flush, String 등등은 일반적인 용법과 의미를 가지고 있습니다.
혼란을 막기 위해, 같은 용법과 의미를 가지고 있는게 아니라면 당신이 만든 method의 이름을
그 이름들 중에서 사용하지 마십시요. 반대로 만약 당신이 구현한 method가 잘 알려진 type과 같은
의미를 가지고 있다면, 같은 이름과 용법을 사용하십시요.
(예: 당신이 만든 string-converter method는 ToString?이 아니라 String이라고 이름지으십시요)



Mixed Caps
마지막으로, Go 언어에서의 규약은 여러 단어를 사용한 이름을 쓰기 위해서는
(_)보다는 대소문자 혼합(MixedCaps?)을 사용한다는 것 입니다.


*/