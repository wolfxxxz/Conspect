package server

import (
	"io"
	"net/http"
)

type StringHandler2 struct {
	message string
}

//Redirect

func (sh StringHandler2) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	Printfln("Request for %v", request.URL.Path)
	switch request.URL.Path {
	case "/favicon.ico":
		http.NotFound(writer, request)
	case "/message":
		io.WriteString(writer, sh.message)
	default:
		http.Redirect(writer, request, "/message", http.StatusTemporaryRedirect)
	}
}

func ServeHTTPAnswer() {
	err := http.ListenAndServe(":5000", StringHandler2{message: "Hello, World"})
	if err != nil {
		Printfln("Error: %v", err.Error())
	}
}

//http://localhost:5000/
