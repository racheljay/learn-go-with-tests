package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// formerly passing in writer with type *bytes.Buffer
func Greet(writer io.Writer, name string) {
	//Fprint is what is called with fmt.Printf
	fmt.Fprintf(writer, "Hello, %s", name)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "worle")
}

func main() {
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler)))
}
