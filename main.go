package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func handler(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadFile("public/index.html") // just pass the file name
    if err != nil {
        fmt.Print(err)
    }
    str := string(b)
    fmt.Fprintf(w, str);
}

func main() {
	http.HandleFunc("/", handler)
	http.Handle("/public/images/", http.StripPrefix("/public/images/", http.FileServer(http.Dir("public/images"))))
	http.ListenAndServe(":8000", nil)
}