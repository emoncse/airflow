package main

import (
	v1 "github.com/emoncse/airflow/api/v1"
	"github.com/emoncse/airflow/configs"
	"log"
	"net/http"
)

func main() {
	// Initialize the database
	configs.InitDB()
	router := v1.RegisterRoutes()
	log.Fatal(http.ListenAndServe(":8080", router))
}
