package servers

import (
	"fmt"
	"net"
)

// Если пришло сообщение кинуть назад ответ
func sendResponse(conn *net.UDPConn, addr *net.UDPAddr) {
	_, err := conn.WriteToUDP([]byte("From server: Hello i got your message "), addr)
	if err != nil {
		fmt.Printf("Couldn't send response %v\n", err)
	}
}

func UDPServer() {

	addr := net.UDPAddr{
		Port: 1234,
		IP:   net.ParseIP("127.0.0.1"),
	}
	ser, err := net.ListenUDP("udp", &addr)

	if err != nil {
		fmt.Printf("Some error %v\n", err)
		return
	}
	fmt.Println("Listening on localhost:1234")

	p := make([]byte, 2048)
	for {
		_, remoteaddr, err := ser.ReadFromUDP(p)

		if err != nil {
			fmt.Printf("Some error %v\n", err)
			continue
		}
		go sendResponse(ser, remoteaddr)
		if len(p) != 0 {
			fmt.Println(string(p))
			break
		}
		//go readMessage(ser, remoteaddr)
	}
}
