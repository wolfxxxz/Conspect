package client

import (
	"bufio"
	"fmt"
	"net/http"
)

func HTTPClientSimpleGet() {
	resp, err := http.Get("http://localhost:8080/hello")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Response status", resp.Status)

	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func HTTPClientHeadersGet() {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://localhost:8080/headers", nil)
	req.Header.Add("test", "test")
	resp, err := client.Do(req)
	//resp,err := http.Get("http:/localhost:8080/headers")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	for name, headers := range resp.Header {
		for _, h := range headers {
			fmt.Printf("H: %v : %v\n", name, h)
		}
	}
	fmt.Println("Response status:", resp.Status)

	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
