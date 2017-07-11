package main

import (
	"log"
	"net/http"

	"github.com/gernest/alien"
)

func main() {
	m := alien.New()
	m.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world"))
	})
	log.Fatal(http.ListenAndServe(":8090", m))
}
