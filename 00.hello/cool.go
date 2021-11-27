package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./")))
	log.Fatal(http.ListenAndServe(":8080", nil))
	// TODO 以后继续开发相关web应用  ——标记
}
