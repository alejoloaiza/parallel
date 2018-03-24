package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	fmt.Println("Hello, playground")
	conn, err := net.Dial("tcp", "130.185.232.126:6667")
	if err != nil {
		conn.Close()
        fmt.Println(err)
	}
	if conn == nil {
		fmt.Println("is null")
	}
	//fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	status, err := bufio.NewReader(conn).ReadString('\n')
	fmt.Println(status)
}

