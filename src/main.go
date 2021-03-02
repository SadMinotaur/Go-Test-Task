package main

import (
	"fmt"
	"log"
	dao "medods/src/dao"
	typesF "medods/src/types"
	"net/http"
	"os"
)

func main() {
	errorLog := log.New(os.Stderr, "ERROR ", log.Ldate|log.Ltime|log.Lshortfile)
	//infoLog := log.New(os.Stdout, "INFO ", log.Ldate|log.Ltime)
	mux := http.NewServeMux()
	//mux.HandleFunc("/first", )
	//mux.HandleFunc("/second", )
	s := dao.SaveToken(typesF.RToken{
		GUID:          "test",
		Token:         "test",
		Created:       0,
		AccessCreated: 0,
		Expires:       0,
		AccessExpires: 0,
	})
	fmt.Print(s)
	server := &http.Server{
		Addr:     ":3210",
		ErrorLog: errorLog,
		Handler:  mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		errorLog.Fatal(err)
	}

}
