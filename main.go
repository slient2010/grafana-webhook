package main

import (
	"fmt"
	// "grafana_webhook/libs"
	"log"
	"net/http"

	"github.com/slient2010/grafana_webhook/libs"
)

func main() {

	addr := ":8080"
	handler := http.DefaultServeMux
	handler.HandleFunc("/sms", libs.HandleWebhook(func(w http.ResponseWriter, b *Body) {

		msg := fmt.Sprintf("Grafana status: %s\n%s", b.Title, b.Message)
		// sendMessage(msg)
		log.Println(msg)

	}, 0))

	go http.ListenAndServe(addr, handler)
	log.Println(fmt.Sprintf("API is listening on: %s", addr))

}
