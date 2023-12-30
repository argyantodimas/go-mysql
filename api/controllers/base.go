package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/argyantodimas/go-mysql/api/models"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type App struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (app *App) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {

	var err error
		DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)
		app.DB, err = gorm.Open(Dbdriver, DBURL)
		if err != nil {
			fmt.Printf("Cannot connect to %s database", Dbdriver)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("We are connected to the %s database", Dbdriver)
		}

	app.DB.Debug().AutoMigrate(&models.User{}, &models.Post{}) //database migration

	app.Router = mux.NewRouter()

	app.initializeRoutes()
}

func (app *App) Run(addr string) {
	fmt.Println("Listening to port 8080")
	log.Fatal(http.ListenAndServe(addr, app.Router))
}
