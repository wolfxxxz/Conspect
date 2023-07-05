package server

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", h, name)
		}
	}
}

func pyramid(w http.ResponseWriter, req *http.Request) {
	var rows int = 10
	temp := 1
	if rows > 0 {
		for i := 0; i < rows; i++ { //
			for j := 2; j <= rows-i; j++ { //пробелы
				fmt.Fprintf(w, " ")
			}
			for k := 0; k <= i; k++ {
				if k == 0 || i == 0 {
					temp = 1
				} else {
					temp = temp * (i - k + 1) / k
				}
				fmt.Fprintf(w, "%d ", temp)
			}
			fmt.Fprintln(w, "")
		}
	}
}

func HTTPServer() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/piramid", pyramid)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}

//insomnia get http://localhost:8080/hello
//answer - hello

//get http://localhost:8080/headers
//answer - insomnia/2023.2.2: User-Agent
//         Accept*
//get http://localhost:8080/got
//answer - 404 page not found

//http://localhost:8080/piramid
