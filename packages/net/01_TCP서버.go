/*
출처: https://coderwall.com/p/wohavg/creating-a-simple-tcp-server-in-go
*/
package main

import (
	"fmt"
	"net"
	"os"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "3333"
	CONN_TYPE = "tcp"
)

func main() {
	//들어오는 접속에 대하여 대기함.
	l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	// 프로그램이 닫힐때 리스너도 끈다.
	defer l.Close()
	fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)
	for {
		// 들어오는 접속을 받는다.
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		// 새로운 Go루틴에서 접속을 다룬다.
		go handleRequest(conn)
	}
}

//들어오는 요청을 다룬다.
func handleRequest(conn net.Conn) {
	// 들어오는 데이터를 보관할 버퍼생성
	buf := make([]byte, 1024)
	// 들어오는 접속을 버퍼로 해서 읽음?
	reqLen, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	// 접속한 사람에게 응답을 돌려보낸다.
	conn.Write([]byte("Message received."))
	// 일을 다 하면 접속을 닫는다.
	conn.Close()
}
