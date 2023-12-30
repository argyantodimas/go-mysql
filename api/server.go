package api

import (
	"fmt"
	"log"
	"os"

	"github.com/argyantodimas/go-mysql/api/controllers"
	"github.com/argyantodimas/go-mysql/api/seed"
	"github.com/joho/godotenv"
)

var app = controllers.App{}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("env file found")
	}
}

func Run() {

	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	app.Initialize(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	seed.Load(app.DB)

	app.Run(":8080")

}
