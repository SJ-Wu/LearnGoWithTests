package di

import (
	"fmt"
	"io"
	"net/http"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func MyGreeterHandler(writer http.ResponseWriter, request *http.Request) {
	Greet(writer, "Chris")
}
