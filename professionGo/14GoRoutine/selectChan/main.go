package main

import (
	"fmt"
	"time"
)

func getInt(ch chan int) {
	for i := 1; i < 5; i++ {
		ch <- i
		time.Sleep(time.Millisecond * 800)
	}
	close(ch)
	fmt.Println("getInt Done")
}

func getStr(ch chan string) {
	ArrString := []string{"hi", "tue", "something", "nothing"}
	for _, strok := range ArrString {
		ch <- strok
		time.Sleep(time.Millisecond * 800)
	}
	close(ch)
	fmt.Println("getString Done")
}

func main() {
	chInt := make(chan int, 2)
	go getInt(chInt)

	chStr := make(chan string, 2)
	go getStr(chStr)

	openCh := 2

	for {
		select {
		case num, ok := <-chInt:
			if ok {
				fmt.Println(num)
			} else {
				fmt.Println("int channel has been closed")
				chInt = nil
				openCh--
			}
		case strok, ok := <-chStr:
			if ok {
				fmt.Println(strok)
			} else {
				fmt.Println("string channel has been closed")
				chStr = nil
				openCh--
			}
		default:
			if openCh == 0 {
				fmt.Println("All channels are closed")
				goto allDone
			}

			time.Sleep(time.Millisecond * 500)
		}
	}
allDone:
	fmt.Println("everithing is the end")
}
