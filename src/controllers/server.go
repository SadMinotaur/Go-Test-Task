package controllers

import (
	"log"
	"net/http"
	"os"
)

type SerLogs struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func Server() (http.Server, *log.Logger, *log.Logger) {
	serLogs := &SerLogs{
		errorLog: log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
		infoLog:  log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/first", serLogs.FirstRoute)
	mux.HandleFunc("/second",  serLogs.SecondRoute)
	return http.Server{
		Addr:     ":3210",
		ErrorLog: serLogs.errorLog,
		Handler:  mux,
	}, serLogs.errorLog, serLogs.infoLog
}
