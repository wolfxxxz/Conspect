package main

import (
	"bufio"
	"fmt"
	"os"
)

func ScanStringOneNewReader() (string, error) {
	fmt.Print("       ...")
	in := bufio.NewScanner(os.Stdin)
	if in.Scan() {
		return in.Text(), nil
	}
	if err := in.Err(); err != nil {
		return "", err
	}
	return "", nil
}
func ScanReaderReadString() (string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введите строку: ")
	//Читает до ограничителя
	input, err := reader.ReadString('7') //if '\n' до ввода
	if err != nil {
		return "", err
	}
	return input, nil

}

func ScanSplit() (string, error) {
	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanLines)
	//in.Split(bufio.ScanWords)

	fmt.Print("Введите строку: ")
	if in.Scan() {
		return in.Text(), nil
	}
	if err := in.Err(); err != nil {
		return "", err
	}
	return "", nil
}
func main() {

}

/*
	str, err := ScanSplit()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(str)
}

/*
	str, err := ScanReaderReadString() //ыаывавы 7 авпва
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(str) //ыаывавы 7
}

/*
		str, err := ScanStringOneNewReader()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(str)
}*/

func Printfln(template string, values ...interface{}) {
	fmt.Printf(template+"\n", values...)
}
