package main

import (
	"gopkg.in/mgo.v2"
	"log"
	"net/http"
)
var mongoConn *mgo.Session
func main() {

	router := NewRouter()

	var err error
	mongoConn, err = createConnection()
	if err != nil {
		panic(err)
	}

	server := http.ListenAndServe(":8000", router)

	log.Fatal(server)
}
