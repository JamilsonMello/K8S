package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/configmap", configMapHandler)

	http.ListenAndServe(":8080", nil)

}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("<h1>Hello World!</h1>"))
}

func configMapHandler(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile(".envfile")

	if err != nil {
		log.Fatalf("Error to reading file: ", err)
	}

	fmt.Fprintf(w, "env file: %s.", string(data))
}
