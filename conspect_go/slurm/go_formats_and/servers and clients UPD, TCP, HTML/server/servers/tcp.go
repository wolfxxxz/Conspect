package servers

import (
	"fmt"
	"net"
	"os"
)

func TCPServer() {
	// Открываем соединение на listen(слушать) или на Write(писать)
	// пакет net сам оформляет (заворачивает пакеты) в зависимости
	// от протокола (tcp - с подтверждением)
	l, err := net.Listen("tcp", "localhost:1234")
	if err != nil {
		fmt.Println("Error listening", err.Error())
		os.Exit(1)
	}

	//если не закрыть приложение то прийдётся искать этот
	//процесс и убивать его принудительно
	//будет висеть в hook (петля) и кушать ресурсы
	defer l.Close()

	//бесконечный цикл который слушает(ждёт) поступление данных
	fmt.Println("Listening on localhost:1234")
	for {
		// метод соединение возвращает net.Conn
		// произошло соединение
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting:", err.Error())
			os.Exit(1)
		}

		// запускаем хендлеры в мультиплекс режиме (паралельно)
		// что бы наш сервер мог принимать несколько соединений
		go handleRequest(conn)
	}

}

func handleRequest(conn net.Conn) {
	//Read []byte{}
	buf := make([]byte, 1024)
	//Байты в буф, количество байт в reqLen
	reqLen, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	fmt.Println(reqLen)
	fmt.Println(string(buf[:reqLen]))

	//Write []byte{}
	bytesWritten, err := conn.Write([]byte("Message received."))
	if err != nil {
		fmt.Println("Error writing:", err.Error())
	}
	fmt.Printf("Bytes written %d\n", bytesWritten)

	//close connection
	conn.Close()
}
