package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()

	message, message_exists := os.LookupEnv("MESSAGE")
	if !message_exists {
		message = "Hello World!"
	}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, message)
	})
	port, port_exists := os.LookupEnv("PORT")
	if !port_exists {
		port = "80"
	}
	log.Println("Запуск веб-сервера на http://127.0.0.1:" + port)
	listener := http.ListenAndServe(":"+port, mux)
	log.Fatal(listener)
}
