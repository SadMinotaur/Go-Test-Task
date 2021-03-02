package main

import (
	serv "medods/src/controllers"
)

func main() {
	server, erLog, _ := serv.Server()
	err := server.ListenAndServe()
	if err != nil {
		erLog.Fatal(err)
	}
}
