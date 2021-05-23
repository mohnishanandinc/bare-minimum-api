package main

import (
	"github.com/bare-minimum-api/cmd/routers"
	"log"
	"net/http"
	"sync/atomic"
	"time"
)

func handleRequests() {
	isReady := &atomic.Value{}
	isReady.Store(false)
	go func() {
		log.Printf("Readyz probe is negative by default...")
		time.Sleep(10 * time.Second)
		isReady.Store(true)
		log.Printf("Readyz probe is positive.")
	}()
	http.HandleFunc("/", routers.HomePage)
	http.HandleFunc("/timestamp", routers.ReturnTimeStamp)
	http.HandleFunc("/healthz", routers.Healthz)
	http.HandleFunc("/readyz", routers.Readyz(isReady))
	log.Fatal(http.ListenAndServe(":8080", nil))
}