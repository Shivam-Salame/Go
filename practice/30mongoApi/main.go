package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/shivamsalame/mongoApi/router"
)

func main() {
	fmt.Println("MongoDB api")
	r := router.Router()
	fmt.Println("server is getting started...")
	log.Fatal(http.ListenAndServe(":5000", r))
	fmt.Println("listening at port :5000 ....")

	// vid-51 troubleshoot
}
