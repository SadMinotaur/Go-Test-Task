package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	errorLog := log.New(os.Stderr, "ERROR ", log.Ldate|log.Ltime|log.Lshortfile)
	//infoLog := log.New(os.Stdout, "INFO ", log.Ldate|log.Ltime)
	mux := http.NewServeMux()
	//mux.HandleFunc("/first", )
	//mux.HandleFunc("/second", )
	server := &http.Server{
		Addr:     ":3210",
		ErrorLog: errorLog,
		Handler:  mux,
	}
	error := server.ListenAndServe()
	if error != nil {
		errorLog.Fatal(error)
	}
}
