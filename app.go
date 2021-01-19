package main

import (
	"fmt"
	"log"
	"net/http"
)

func handlePostForm(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	username := request.Form.Get("username")
	password := request.Form.Get("password")

	fmt.Printf("username=%s, password=%s\n", username, password)
	fmt.Fprintf(writer, `{"code":2222}`)
	//return "{code:0}"
}
func main() {
	http.HandleFunc("/handlePostForm", handlePostForm)
	log.Println("Running at port 9999 ...")
	err := http.ListenAndServe("0.0.0.0:9999", nil)
	if err != nil {
		log.Fatal("ListenAndServe:`{\"code\":0}` ", err.Error())
	}
}
