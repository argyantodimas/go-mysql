package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/argyantodimas/go-mysql/api/auth"
	"github.com/argyantodimas/go-mysql/api/models"
	"github.com/argyantodimas/go-mysql/api/responses"
	"github.com/argyantodimas/go-mysql/api/utils"
	"golang.org/x/crypto/bcrypt"
)

func (app *App) Login(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user.Prepare()
	err = user.Validate("login")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	token, err := app.SignIn(user.Email, user.Password)
	if err != nil {
		formattedError := utils.FormatError(err.Error())
		responses.ERROR(w, http.StatusUnprocessableEntity, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, token)
}

func (app *App) SignIn(email, password string) (string, error) {

	var err error

	user := models.User{}

	err = app.DB.Debug().Model(models.User{}).Where("email = ?", email).Take(&user).Error
	if err != nil {
		return "", err
	}
	err = models.VerifyPassword(user.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}
	return auth.CreateToken(user.ID)
}
