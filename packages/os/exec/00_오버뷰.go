// exec 패키지는 외부 커맨트를 실행시킨다.
// 이것은 os.StartProcess를 포장해서 stdin 과 stdour를 remap하기, pipe로 I/O에 접속하기 그리고 다른 수정등을 하기 쉽게 해준다.
//

//Package exec runs external commands. It wraps os.StartProcess to make it easier to remap stdin and stdout,
//connect I/O with pipes, and do other adjustments.