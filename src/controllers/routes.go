package controllers

import (
	"log"
	"net/http"
	"os"
)

func Server() (http.Server, *log.Logger, *log.Logger) {
	errorLog := log.New(os.Stderr, "ERROR ", log.Ldate|log.Ltime|log.Lshortfile)
	infoLog := log.New(os.Stdout, "INFO ", log.Ldate|log.Ltime)
	mux := http.NewServeMux()
	mux.HandleFunc("/first", FirstRoute)
	mux.HandleFunc("/second", SecondRoute)
	return http.Server{
		Addr:     ":3210",
		ErrorLog: errorLog,
		Handler:  mux,
	}, errorLog, infoLog
}
