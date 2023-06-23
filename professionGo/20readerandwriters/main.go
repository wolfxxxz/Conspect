package main

import (
	"fmt"
	"io"
	"os"
)

func processData(reader io.Reader, writer io.Writer) {
	count, err := io.Copy(writer, reader)
	if err == nil {
		Printfln("Read %v bytes", count)
	} else {
		Printfln("Error: %v", err.Error())
	}
}

// os.Readfile(pathfile)
func OsReadfile(path string) {
	dat, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(dat))
}

// Чтение с проверкой наличия и прав доступа
// С использованием буфера
func OsOpenAndRead(filepath string) {
	f, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err)
	}
	b1 := make([]byte, 32)
	n1, err := f.Read(b1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b1[:n1]))
	f.Close()
}

func main() {
	OsReadfile("text.txt")
	OsOpenAndRead("text.txt")
}

/*
func main() {

	var w1 strings.Builder
	var w2 strings.Builder
	var w3 strings.Builder
	combinedWriter := io.MultiWriter(&w1, &w2, &w3)
	//Тут рождается каяк
	GenerateData(combinedWriter)
	Printfln("Writer #1: %v", w1.String())
	Printfln("Writer #2: %v", w2.String())
	Printfln("Writer #3: %v", w3.String())

}*/
