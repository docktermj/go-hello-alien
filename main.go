package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gernest/alien"
)

func main() {
	m := alien.New()

	// Hello world.

	m.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world"))
	})

	// Streaming.

	m.Get("/streaming", func(w http.ResponseWriter, r *http.Request) {

		// Need to cast http.ResponseWriter as a "flusher".

		flusher, ok := w.(http.Flusher)
		if !ok {
			panic("expected http.ResponseWriter to be an http.Flusher")
		}

		// Stream the response body.  Note: loop will not stop.

		ticker := time.NewTicker(time.Millisecond * 500)
		for aTime := range ticker.C {
			w.Write([]byte(aTime.String()))
			flusher.Flush()
		}
	})

	log.Fatal(http.ListenAndServe(":8090", m))
}
