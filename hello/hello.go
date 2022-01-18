package hello

import (
	"fmt"
	"io"
)

func Hello() string {
	return "Hello, world"
}

func Greet(writer io.Writer, name string) {
	// fmt.Printf("Hello, %s", name)
	fmt.Fprintf(writer, "Hello, %s", name)
}
