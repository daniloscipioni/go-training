package main

import (
	"fmt"
	"http-server/users"
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

func usersPrint(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, users.PrintUsers("Agora foi essa bagaça"))
}

func main() {

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/users", usersPrint)

	http.ListenAndServe(":8090", nil)
}
