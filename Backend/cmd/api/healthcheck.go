package main

import (
	"net/http"
)

func (app *Application) healthcheck (w http.ResponseWriter, r *http.Request){
    status := map[string]string{
        "status": "active",
        "environment": app.config.env,
        "version": Version,
    }
    err := app.writeJSON(w, http.StatusOK, status, nil)
    if err != nil {
        http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
    }
}
