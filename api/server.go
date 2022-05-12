package api

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/mikorail/go-api/api/controllers"
)

var server = controllers.Server{}

func Run() {
	var err error
	err = godotenv.Load()

	if err != nil {
		log.Fatalf(".env can't be processed %v", err)
	} else {
		fmt.Println(".env successfully processed")
	}

	server.Initialize(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	server.Run(":8090")
}
