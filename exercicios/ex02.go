package main

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
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}

}

func main1() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	//http.HandleFunc("/users", users.printUsers("Danilo"))

	http.ListenAndServe(":8090", nil)
}
