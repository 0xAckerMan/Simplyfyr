package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/0xAckerMan/Simplifyr/Backend/internal/data"
)

func (app *Application) all_roles(w http.ResponseWriter, r *http.Request) {
	var roles []*data.Role
	roles, err := app.models.Roles.Get_all()
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"roles": roles}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *Application) single_role(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDparam(r)
	if err != nil {
		app.notFoundResponse(w, r)
	}
	role, err := app.models.Roles.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"role": role}, nil)
}

func (app *Application) create_role(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Role string `json:"role"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	role := &data.Role{
		Role: input.Role,
	}

	err = app.models.Roles.Insert(role)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	header := make(http.Header)
	header.Set("Location", fmt.Sprintf("/v1/roles/%d", role.Id))

	err = app.writeJSON(w, http.StatusOK, envelope{"role": role}, header)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *Application) update_role(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDparam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	role, err := app.models.Roles.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
	}

    var input struct{
        Role *string `json:"role"`
    }

    err = app.readJSON(w,r,&input)
    if err != nil {
        app.badRequestResponse(w,r,err)
        return
    }

    if input.Role != nil{
        role.Role = *input.Role
    }

    err = app.models.Roles.Update(role)
    if err != nil{
        app.serverErrorResponse(w,r,err)
        return
    }

    app.writeJSON(w,http.StatusOK,envelope{"role":role}, nil)

}

func (app *Application) delete_role(w http.ResponseWriter, r *http.Request){
    id, err := app.readIDparam(r)
    if err != nil{
        app.notFoundResponse(w, r)
        return
    }
    err = app.models.Roles.Delete(id)
    if err != nil{
        switch{
        case errors.Is(err, data.ErrRecordNotFound):
        app.notFoundResponse(w,r)
        default:
        app.serverErrorResponse(w,r,err)
    }
        return
    }

    err = app.writeJSON(w, http.StatusOK,envelope{"message": "Role deleted successfully"}, nil)
    if err != nil{
        app.serverErrorResponse(w,r,err)
        return
    }

}
