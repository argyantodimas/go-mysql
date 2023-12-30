package controllers

import (
	"net/http"

	"github.com/argyantodimas/go-mysql/api/responses"
)

func (app *App) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome to API")

}
