package main

import (
	"io"
	"log"
	"net"
)

func main() {
	// 모든 인터페이스들 위에서 2000포트로 TCP 접속을 대기함.
	l, err := net.Listen("tcp", ":2000")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	for {
		//접속대기
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		// 새로운 Go루틴에서 접속을 다룸.
		// loop 는 접속으로 돌아간다, 그래서 여러개의 접속이 동시적으로 처리된다.
		go func(c net.Conn) {
			// 들어오는 모든 데이터를 출력함.
			io.Copy(c, c)
			//접속 종료
			c.Close()
		}(conn)
	}
}
