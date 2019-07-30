package main

import (
	"fmt"
	"log"
	"net/http"

	ce "github.com/cloudevents/sdk-go"
	ev "github.com/mchmarny/gcputil/env"
)

func main() {

	port := ev.MustGetIntEnvVar("PORT", 8080)

	// Handler Mux
	mux := http.NewServeMux()

	// Ingres API Handler
	t, err := ce.NewHTTPTransport(
		ce.WithMethod("POST"),
		ce.WithPath("/"),
		ce.WithPort(port),
	)
	if err != nil {
		log.Fatalf("failed to create CloudEvents transport, %s", err.Error())
	}

	// wire handler for CE
	t.SetReceiver(&eventReceiver{})

	// Health Handler
	mux.HandleFunc("/health", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, "ok")
	})

	// Events or UI Handlers
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Method, %s", r.Method)
		if r.Method == "POST" {
			t.ServeHTTP(w, r)
			return
		}
		fmt.Fprint(w, "Nothing to see here. Use POST to send CloudEvents")
	})

	a := fmt.Sprintf(":%d", port)
	log.Fatal(http.ListenAndServe(a, mux))

}
