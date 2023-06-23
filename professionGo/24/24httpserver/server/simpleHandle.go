package server

import (
	"io"
	"net/http"
)

type StringHandler3 struct {
	message string
}

//Redirect

func (sh StringHandler3) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	Printfln("Request for %v", request.URL.Path)
	io.WriteString(writer, sh.message)
}

func ServeHTTPHandle() {
	http.Handle("/message", StringHandler{"Hello, World"})
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.Handle("/", http.RedirectHandler("/message", http.StatusTemporaryRedirect))

	err := http.ListenAndServe(":5000", StringHandler2{message: "Hello, World"})
	if err != nil {
		Printfln("Error: %v", err.Error())
	}
}

//http://localhost:5000/
