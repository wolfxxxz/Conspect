package main

import "fmt"

func Printfln(template string, values ...interface{}) {
	fmt.Printf(template+"\n", values...)
}
func main() {
	//Printfln("Value: %v", Kayak)
	//Printfln("Value with fields: %+v", Kayak)
	//ScanArr()
	//SSscan()
	SSscan2()
}
