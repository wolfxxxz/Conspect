package main

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
	reqLen, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	fmt.Println(reqLen)

	//Write []byte{}
	bytesWritten, err := conn.Write([]byte("Message received."))
	if err != nil {
		fmt.Println("Error writing:", err.Error())
	}
	fmt.Printf("Bytes written %d\n", bytesWritten)

	//close connection
	conn.Close()
}

/*
//После прочтения символов флажок смещается на
//количество прочитанных символов

func read(filepath string) {
	//Получаем файл
	f, err := os.Open(filepath)
	check(err)

	//Вычитываем первые 7 байт
	b2 := make([]byte, 7)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%s\n", string(b2[:n2])) //Hello

	// Читаем Следующие 7 байт
	b4 := make([]byte, 7)
	n4, err := f.Read(b4)
	check(err)
	fmt.Println("n4 ", n4)         //n4  7
	fmt.Printf("%s\n", string(b4)) //how are

	// Смещение позиции чтения
	//сместить на количество байт
	//сместить можно вперёд или назад + -
	//offset := int64(7)
	offset := int64(-7)
	//откуда - (0 - с начала, 1 - с текущей позиции, с конца)
	whence := 2
	_, err = f.Seek(offset, whence)
	check(err)

	// Читаем следующие 7 байт
	b3 := make([]byte, 7)
	n3, err := f.Read(b3)
	check(err)
	fmt.Printf("%s\n", string(b3[:n3])) //Hello

	f.Close()
}

func main() {
	read("test.txt")
	//Hello
	//how are you
	//i'm well
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

/*
func main() {
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	delimiter := byte('\n') // Определяем символ-разделитель
	arrStr := []string{}

	for {
		line, err := reader.ReadString(delimiter)
		if err != nil {
			// Если конец файла
			break
		}

		// Убираем символ-разделитель из строки
		line = strings.TrimSuffix(line, string(delimiter))

		arrStr = append(arrStr, string(line))
		//fmt.Println("Read line:", line)

		// Если достигнут конец файла, прекращаем чтение
		if err == bufio.ErrBufferFull || err == io.EOF {
			break
		}
	}
	fmt.Println(arrStr)
	fmt.Println("the end")
}

/*
func read(filepath string) {
	//проверяет наличие и доступность прав на чтение файла
	//если всё ок возвращает открытый файл
	f, err := os.Open(filepath)
	check(err)

	//os.Read([]byte)
	//Читает байты (с открытого файла) сколько влезет в массив
	b1 := make([]byte, 7)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1))

	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%v\n", string(b2[:n2]))

	// Выполнение операций с открытым файлом
	// Изменение позиции чтения/записи
	// offset := int64(6)
	// whence := 0
	// newPos, err := f.Seek(offset, whence)
	o3, err := f.Seek(6, 0)
	check(err)

	// Считает только количество байт
	// f - источник (os.Open)
	// b3 - буфер в который будут прочитаны данные
	// 2 - гарантированное количество прочитанных байт
	// n3 - количество байтов, фактически
	//прочитанных из источника данных.
	// err - mistake
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("io.ReadAtLeast:  %d bytes @ %d: %s\n", n3, o3, string(b3))
	_, err = f.Seek(0, 0)
	check(err)

	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(15)
	check(err)
	fmt.Printf("15 bytes: %s\n", string(b4))

	f.Close()
}

func main() {
	read("test.txt")
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
*/
