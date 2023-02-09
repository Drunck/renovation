package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"s_renovation.net/internal/data"
	"s_renovation.net/validator"
	"time"
)

func (app *application) registerUserHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name     string `form:"name"`
		Surname  string `form:"surname"`
		Email    string `form:"email"`
		Password string `form:"password"`
	}
	err := app.decodePostForm(r, &input)
	//err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		app.errorLog.Print("server error response")
		app.errorLog.Println(err)
	}
	user := &data.User{
		Name:      input.Name,
		CreatedAt: time.Now(),
		Surname:   input.Surname,
		Email:     input.Email,
	}
	err = user.Password.Set(input.Password, 12)
	if err != nil {
		app.errorLog.Print("server error response")
		app.errorLog.Print(err)
	}

	v := validator.New()
	if data.ValidateUser(v, user); !v.Valid() {
		app.errorLog.Print("failed to validate user")
		app.errorLog.Println(v.Errors)
	}
	err = app.models.User.Insert(user)
	if err != nil {
		app.errorLog.Print("server error response")
		app.errorLog.Print(err)
	}

}

func (app *application) ShowUserHandler(w http.ResponseWriter, r *http.Request) {
	var email string
	err := json.NewDecoder(r.Body).Decode(&email)
	if err != nil {
		app.errorLog.Print("server error response")
	}
	v := validator.New()
	if data.ValidateEmail(v, email); !v.Valid() {
		app.errorLog.Print("couldn't validate email")
		app.errorLog.Println(v.Errors)
	}
	user, err := app.models.User.GetByEmail(email)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.infoLog.Print("user not found")

		default:
			app.errorLog.Print("server error response")
			app.errorLog.Print(err)
		}
	}
	w.Write([]byte(user.Name))
}
