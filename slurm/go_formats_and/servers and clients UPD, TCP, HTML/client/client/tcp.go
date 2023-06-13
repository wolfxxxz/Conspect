package client

import (
	"bufio"
	"fmt"
	"net"
)

func TCPClient() {
	p := make([]byte, 1024)
	conn, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		fmt.Printf("Some err #{err}")
		return
	}
	fmt.Fprintf(conn, "H1 TCP Server, how are you doing?")
	_, err = bufio.NewReader(conn).Read(p)

	if err == nil {
		fmt.Printf("%s", p)
	} else {
		fmt.Printf("Some error #{err}\n")
	}
	conn.Close()
}
