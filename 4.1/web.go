package main

import (
	"fmt"
	"log"
	"html/template"
	"net/http"
)

func sayHello(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "hello, go")
}

func login(writer http.ResponseWriter, request *http.Request) {
	method := request.Method
	fmt.Println("method : ", method)
	if method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(writer, nil)		
	} else {
		request.ParseForm()
		fmt.Println("username : ", request.Form["username"])
		fmt.Println("password : ", request.Form["password"])
		template.HTMLEscape(writer, []byte(request.Form.Get("username")))
	}
}

func main() {
	http.HandleFunc("/", sayHello)
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe : ", err)
	}
}
