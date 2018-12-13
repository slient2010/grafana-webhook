package main

import (
	"fmt"
	"gwebhook/libs"
	"log"
	"net/http"
)

func main() {

	addr := ":9000"
	handler := http.DefaultServeMux
	handler.HandleFunc("/sms", libs.HandleWebhook(func(w http.ResponseWriter, b *libs.Body) {

		msg := fmt.Sprintf("Grafana status: %s\n%s\n %s", b.Title, b.Message, b.State)
		// sendMessage(msg)
		log.Println(msg)

	}, 0))

	log.Println(fmt.Sprintf("API is listening on: %s", addr))

	// http.ListenAndServe(addr, handler)

	log.Fatal(http.ListenAndServe(addr, handler))
}
