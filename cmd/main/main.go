package main

import (
	"log"
	"net/http"
	"user-authentication-with-go/pkg/config"
	"user-authentication-with-go/pkg/routes"
)

func main() {
	db := config.InitDB()
	r := routes.SetupRoutes(db)
	log.Fatal(http.ListenAndServe(":8000", r))
}
