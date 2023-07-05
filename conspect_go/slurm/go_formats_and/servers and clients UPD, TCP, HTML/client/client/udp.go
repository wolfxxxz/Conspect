package client

import (
	"bufio"
	"fmt"
	"net"
)

func UDPClient() {
	p := make([]byte, 1024)
	conn, err := net.Dial("udp", "127.0.0.1:1234")
	if err != nil {
		fmt.Printf("Some err %v\n", err)
		return
	}
	fmt.Fprintf(conn, "H1 UDP Server, how are you doing?")
	_, err = bufio.NewReader(conn).Read(p)

	if err == nil {
		fmt.Printf("%s\n", p)
	} else {
		fmt.Printf("Some err %v\n", err)
	}
	conn.Close()
}
