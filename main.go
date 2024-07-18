package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	err := http.ListenAndServe("0.0.0.0:"+port, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("hi golang"))
		if err != nil {
			panic(err)
		}
	}))
	if err != nil {
		panic(err)
	}
}
