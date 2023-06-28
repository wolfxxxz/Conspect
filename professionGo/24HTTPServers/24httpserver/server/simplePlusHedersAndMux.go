package server

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func hello1(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func headers1(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", h, name)
		}
	}
}

func pyramid1(w http.ResponseWriter, req *http.Request) {
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

func pyramidPlusId(w http.ResponseWriter, req *http.Request) {
	rows, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		fmt.Println("piramid err", err)
	}

	temp := 1
	/*
		temp := "Server are working"
		for i := 0; i < rows; i++ {
			for k := 0; k <= i; k++ {
				fmt.Fprintf(w, "%v ", temp)
			}
			fmt.Fprintln(w, "")
		}*/

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

func HTTPServerMux() {
	router := mux.NewRouter()
	router.HandleFunc("/hello", hello1)
	router.HandleFunc("/headers", headers1)
	router.HandleFunc("/piramid", pyramid1)
	router.HandleFunc("/piramid"+"/{id}", pyramidPlusId)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println(err)
	}
}

//insomnia get http://localhost:8080/hello
//answer - hello

//get http://localhost:8080/headers
//answer - insomnia/2023.2.2: User-Agent
//         */*: Accept
//get http://localhost:8080/got
//answer - 404 page not found

//http://localhost:8080/piramid

//http://localhost:8080/piramid/5
