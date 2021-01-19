package main

import (
	"io"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello,world")

}
func main() {
	http.HandleFunc("/hello", helloHandler)
	err := http.ListenAndServe(":8066", nil)
	if err != nil {
		log.Fatal("ListenandServer:", err.Error())
	}
}
