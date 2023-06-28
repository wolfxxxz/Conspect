package server

import (
	"fmt"
	"io"
	"net/http"
)

func Printfln(template string, values ...interface{}) {
	fmt.Printf(template+"\n", values...)
}

type StringHandler1 struct {
	message string
}

func (sh StringHandler1) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path == "/favicon.ico" {
		Printfln("Request for icon detected - returning 404")
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	Printfln("Request for %v", request.URL.Path)
	io.WriteString(writer, sh.message)
}

func ServeHTTPFavIso() {
	err := http.ListenAndServe(":5000", StringHandler1{message: "Hello, World"})
	if err != nil {
		Printfln("Error: %v", err.Error())
	}
}

//http://localhost:5000/
