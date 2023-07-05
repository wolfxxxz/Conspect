package main

import "server/client"

func main() {

	//client.TCPClient()

	//client.UDPClient()

	client.HTTPClientSimpleGet()
	client.HTTPClientHeadersGet()

}
