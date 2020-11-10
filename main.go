package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"webpay-rest-go/routers"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%s", os.Getenv("PORT")),
		Handler:        router,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()

}
