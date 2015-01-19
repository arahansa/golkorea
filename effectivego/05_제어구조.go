// 05_제어구조.go
// https://code.google.com/p/golang-korea/wiki/EffectiveGo#제어구조(Control_structures)

/*

   Go언어의 제어문은 C와 연관성이 있지만 중요한 차이가 있습니다.
   Go언어의 제어문에는 do와 while 루프가 없습니다.
   오직 좀 더 일반화된 for 루프와 switch문이 있으며 더 탄력적입니다.

   if와 switch문은 for루프와 같이 추가적인 초기화를 지원하며,
   type switch와 multiway communications multiplexer, select와 같은 추가적인 새로운 제어문이 있습니다.
   문법 또한 조금 다릅니다. 괄호()는 필요하지 않으며 body는 항상 중괄호({})로 구분되어져야 합니다.

   If
   Go 언어에서 if문의 간단한 예는 다음과 같습니다
*/
if x > 0 {
	return y
}
/*
   중괄호({})의 의무적인 사용은 다중 라인에서 if 문을 단순하게 만들어줍니다.
   특히, body안에 return이나 break같은 제어문을 포함하고 있는 경우에는 괜찮은 방법입니다.

   if나 switch 문이 초기화를 지원하기 때문에, if나 switch문에서 지역변수를 설정하기 위해
   이 방법을 사용하는 것을 쉽게 발견할 수 있습니다.
*/
if err := file.Chmod(0664); err != nil {
	log.Print(err)
	return err
}
/*
   Go언어의 라이브러리에서 if문이 다음 명령문으로 진행되지 않는 경우
   (break, continue, goto 또는 return 때문에 body가 종료되는 상황)
   불필요한 else는 생략되는 것을 발견하게 될 것 입니다.
*/

f, err := os.Open(name, os.O_RDONLY, 0)
if err != nil {
	return err
}
codeUsing(f)

/*
   다음은 일련의 error에 대응하기 위해 code가 보호되어야 하는 일반적인 상황에 대한 예제입니다.
   만약 제어의 흐름이 성공적이라면 그 코드는 잘 동작할 것이고, error가 발생할 때마다
   error를 제거할 것입니다. error는 return문에서 발생하는 경향이 있기 때문에 결과 code는 else문이 필요 없습니다.
*/

   f, err := os.Open(name, os.O_RDONLY, 0)
   if err != nil {
       return err
   }
   d, err := f.Stat()
   if err != nil {
       return err
   }
   codeUsing(f, d)
/*
   For
   Go언어에서 for 루프는 C언어의 for 루프와 비슷하지만 같지는 않습니다. 
for문과 while문은 구분되어지며 do-while문은 없습니다. 
세가지 형태가 있으며 그중 하나만 세미콜론을 가지고 있습니다.
*/

   // C언어와 유사한 형태의 for문
   for init; condition; post { }

   // C언어와 유사한 형태의 while문
   for condition { }

   // C언어와 유사한 형태의 for(;;)
   for { }
  
   //루프에서는 짧은 선언이 index 변수를 선언하기 쉽게 만듭니다.

   sum := 0
   for i := 0; i < 10; i++ {
       sum += i
   }
   //만약 당신이 배열,slice, string, map, reading을 넘어서 루프를 돌고 있다면, 
   //range문이 당신을 대신해서 루프를 관리해줄 것 입니다.

   var m map[string]int
   sum := 0
   for _, value := range m {  // key is unused
       sum += value
   }
   //string의 경우, range문은 UTF-8 파싱에 의한 개별적인 유니코드 문자를 해결하는데 더욱 유익할 것입니다. 
   //잘못된 encoding은 하나의 byte를 제거하고 U+FFFD 문자로 대체할 것입니다.

   //아래 루프는

   for pos, char := range "日本語" {
       fmt.Printf("character %c starts at byte position %d\n", char, pos)
   }
   //다음과 같이 출력됩니다.

   character 日 starts at byte position 0
   character 本 starts at byte position 3
   character 語 starts at byte position 6
   //마지막으로 Go언어는 콤마(,) 오퍼레이터가 없으며 ++, --는 표현식이 아니라 명령문입니다. 
   //그러므로 만약 for문 안에서 다중 변수(multiple variable)를 사용하려 한다면 병렬 할당(parallel assignment)을 사용해야만 합니다.

   // Reverse a
   for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
       a[i], a[j] = a[j], a[i]
   }
   Switch
   Go언어에서 switch문은 C언어에서 보다 더 일반적입니다. 표현식은 상수이거나 integer일 필요가 없습니다. case문은 맨위에서 밑에까지 맞는 값을 찾을 때까지 계속 값을 비교할 것입니다. 그리고 switch문에서는 switch문을 true로 하는 표현식이 없기 때문에 switch문 대신 if-else-if-else 체인을 사용하는 것도 가능합니다.

   func unhex(c byte) byte {
       switch {
       case '0' <= c && c <= '9':
           return c - '0'
       case 'a' <= c && c <= 'f':
           return c - 'a' + 10
       case 'A' <= c && c <= 'F':
           return c - 'A' + 10
       }
       return 0
   }
   Go언어에서 switch문의 경우 자동으로 통과하는 방법을(case문을 통과하는 경우를 말함) 사용할 수 없지만 콤마(,)를 사용하여 아래 예문과 같이 각 case를 구분할 수 있습니다.

   func shouldEscape(c byte) bool {
       switch c {
       case ' ', '?', '&', '=', '#', '+', '%':
           return true
       }
       return false
   }
   여기에 두개의 switch문을 사용하여 byte 배열을 비교하는 루틴의 예가 있습니다.

   // Compare 함수는 두 byte의 배열들을 비교하여 하나의 integer값을 리턴해준다.
   // 결과값은 만약 a == b라면 0, a < b라면 -1, a > b라면 +1을 리턴할 것이다.
   func Compare(a, b []byte) int {
       for i := 0; i < len(a) && i < len(b); i++ {
           switch {
           case a[i] > b[i]:
               return 1
           case a[i] < b[i]:
               return -1
           }
       }
       switch {
       case len(a) < len(b):
           return -1
       case len(a) > len(b):
           return 1
       }
       return 0
   }
   또한, switch문은 동적인 형태의 interface 변수를 찾는데 사용할 수 있습니다. 그런 형태의 switch문은 괄호() 안에서 키워드의 type과 assertion type의 문법을 사용할 수 있습니다. 만약 switch문이 expression에서 변수를 선언했다면, 변수는 각 문장에 대응하는 type을 가지게 될 것입니다.

   switch t := interfaceValue.(type) {
   default:
       fmt.Printf("unexpected type %T", t)  // %T prints type
   case bool:
       fmt.Printf("boolean %t\n", t)
   case int:
       fmt.Printf("integer %d\n", t)
   case *bool:
       fmt.Printf("pointer to boolean %t\n", *t)
   case *int:
       fmt.Printf("pointer to integer %d\n", *t)
   }

*/