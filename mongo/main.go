package main

import (
	"fmt"
	"log"
	"mondoSetup/routers"
	"net/http"
)

func main() {
	fmt.Println("MongoDB Setup")
	r := routers.Router()
	fmt.Println("Server Started")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening on Port 4000")

}
