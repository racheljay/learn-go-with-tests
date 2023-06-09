package main

import (
	"bytes"
	"fmt"
)

func Greet(writer *bytes.Buffer, name string) {
	//Fprint is what is called with fmt.Printf
	fmt.Fprintf(writer, "Hello, %s", name)
}
