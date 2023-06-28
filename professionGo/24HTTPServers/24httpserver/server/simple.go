package server

import (
	"fmt"
	"io"
	"net/http"
)

type StringHandler struct {
	message string
}

// Типа есть метод который принимает  (writer http.ResponseWriter, request *http.Request)
// Значит StringHandler удовлетворяет
// интерфейсу func http.ListenAndServe(addr string, handler http.Handler) error
// http.Handler

func (sh StringHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	io.WriteString(writer, sh.message)
}

func ServerSimple() {
	//а значит функция может принимать StringHandler
	err := http.ListenAndServe(":5000", StringHandler{message: "Hello, World"})
	if err != nil {
		fmt.Printf("Error: %v\n", err.Error())
	}
}
