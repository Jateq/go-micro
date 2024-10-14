package main

import (
	"fmt"
	"log"
	"net/http"
)

const webPort = "80"

type Config struct {}

func main() {
	app := Config{}
	log.Println("Starting broker service on port %s", webPort)

	srv := &http.Server{
		Addr: fmt.Sprintf(":%s", webPort),
	}

	err := srv.ListenAndServe()
	if err != nil{
		log.Panic(err)
	}
}