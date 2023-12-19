package main

import (
	"errors"
	"net/http"

	"github.com/0xAckerMan/Simplifyr/Backend/internal/data"
)

func (app *Application) registerUserHandler(w http.ResponseWriter, r *http.Request){
    var input struct{
        Name string `json:"name"`
        Email string `json:"email"`
        Password string `json:"password"`
    }

    err := app.readJSON(w,r,&input)
    if err != nil{
        app.badRequestResponse(w,r,err)
        return
    }
    user := data.User{
        Name: input.Name,
        Email: input.Email,
    }
    err = user.Password.Set(input.Password)
    if err != nil{
        app.serverErrorResponse(w,r,err)
        return
    }

    err = app.models.Users.Insert(&user)
    if err != nil {
        switch{
        case errors.Is(err, data.ErrDuplicateEmail):
        app.errorResponse(w, r, http.StatusBadRequest, "the email already exists")
        default:
        app.serverErrorResponse(w,r, err)
    }
        return
    }

    err = app.writeJSON(w, http.StatusCreated, envelope{"user": user}, nil)
    if err != nil{
        app.serverErrorResponse(w, r, err)
        return
    }

}
