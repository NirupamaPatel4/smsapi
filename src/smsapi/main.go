package main

import (
	"github.com/valyala/fasthttp"
	"log"
	"smsapi/router"
)

func main() {
	addr := ":8080"
	log.Printf("Starting HTTP server on %q", addr)
	go func() {
		if err := fasthttp.ListenAndServe(addr, router.HandleRoute); err != nil {
			log.Fatalf("error starting server: %s", err)
		}
	}()
	select {}
}
