package main

import (
	"log"
	"net/http"
	"net/url"
	"os"
)

func main() {

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	bindAddr := os.Getenv("BIND")
	if bindAddr == "" {
		bindAddr = ":7000"
	}

	log.Println("Starting server on " + bindAddr)

	err := http.ListenAndServe(bindAddr, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		originUrl := r.URL.Query().Get("mirror")
		path, err := url.Parse(originUrl)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Invalid mirror url"))
			return
		}
		proxy := NewReverseProxy(path)
		proxy.ServeHTTP(w, r)
	}))

	panic(err)
}
