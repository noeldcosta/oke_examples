package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"net/http"
	"os"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello JD Edwards - test 1b !!"))
}

func main() {
	fmt.Printf("Starting Hello JDE application... \n")
	http.HandleFunc("/", HelloWorld)
	http.ListenAndServe(":8080", handlers.LoggingHandler(os.Stdout, http.DefaultServeMux))
}
