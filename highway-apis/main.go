package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello, World!")

	// create handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Hello, World!")
		d, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Println("Data : ", string(d))
	})
	// another Handler with Path
	http.HandleFunc("/goodbye", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Goodbye, World!")
	})
	// create server
	http.ListenAndServe(":8080", nil)

}
