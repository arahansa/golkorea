//http://golang.org/pkg/fmt/
/*
fmt 패키지는 형식화된 입출력을 c의 printf, scanf와 비슷한 함수들로 구현한다.
형식 verb들은 C에서 따왔지만, 좀 다 간결하다.

*/
%v	구조체를 출력할 때 기본형식값, +표시(%+v)는 필드명을 더한다.
%v	the value in a default format when printing structs, the plus flag (%+v) adds field names
%#v	값을 재표시하는 Go문법
%T	값의 타입을 재표시(representation)
%%	a literal percent sign; consumes no value

Boolean:

%t	참, 거짓

Integer:

%b	2진수
%c	문자, 유니코드에 해당하는, 
%d	10진수
%o	8진수
%q	단일인용문자 리터럴? a single-quoted character literal safely escaped with Go syntax.
%x	16진수 영어는 소문자 표현 //base 16, with lower-case letters for a-f
%X	16진수 대문자로 표현 //base 16, with upper-case letters for A-F
%U	유니코드 포맷 //Unicode format: U+1234; same as "U+%04X"

//TODO 

