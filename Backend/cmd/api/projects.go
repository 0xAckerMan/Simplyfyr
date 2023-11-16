package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/0xAckerMan/Simplifyr/Backend/internal/data"
)

func (app *Application) all_projects(w http.ResponseWriter, r *http.Request) {
	message := "All projects lists"
	err := app.writeJSON(w, http.StatusOK, message, nil)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

func (app *Application) single_project(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDparam(r)
	if err != nil{
		http.NotFound(w, r)
	}

	project := data.Project{
		Id:           id,
		Name:         "Simplifyr",
		Category:     []string{"Web", "API", "ALX"},
		Excerpt:      "A simple web api",
		Description:  "A simple web api for a project management system",
		Assigned_to:  2,
		Created_by:   1,
		Created_date: time.Now(),
		Due_date:     time.Now().Add(72 * time.Hour),
		Done:         0,
	}
	err = app.writeJSON(w, http.StatusOK, project, nil)
}

func (app *Application) pm_projects(w http.ResponseWriter, r *http.Request) {
    id, err := app.readIDparam(r)
    if err != nil{
        http.NotFound(w,r)
    }
    fmt.Fprintf(w, "all projects belonging to pmof id %d \n", id)
}

func (app *Application) employee_projects (w http.ResponseWriter, r *http.Request){
    id, err := app.readIDparam(r)
    if err != nil{
        http.NotFound(w,r)
    }

    fmt.Fprintf(w, "all employee projects of id %d \n", id)
}

func (app *Application) employee_completed_projects(w http.ResponseWriter, r *http.Request){
    id,err := app.readIDparam(r)
    if err != nil {
        http.NotFound(w, r)
    }
    fmt.Fprintf(w, "Employee of id %d completed projects \n", id)
}

func (app *Application) mark_complete(w http.ResponseWriter, r *http.Request){
    id, err := app.readIDparam(r)
    if err != nil{
        http.NotFound(w, r)
    }
    fmt.Fprintf(w, "project %d marked as completed \n", id)
}

func (app *Application) mark_uncomplete(w http.ResponseWriter, r *http.Request){
    id,err := app.readIDparam(r)
    if err != nil {
        http.NotFound(w, r)
        return
    }
    fmt.Fprintf(w, "project of id %d marked uncomplete \n", id)
}

func (app *Application) create_project(w http.ResponseWriter, r *http.Request){
    fmt.Fprintln(w, "created a new project")
}

func (app *Application) update_project (w http.ResponseWriter, r *http.Request){
    id, err := app.readIDparam(r)
    if err != nil{
        http.NotFound(w, r)
    }
    fmt.Fprintf(w, "update project of id %d \n", id)
}

func (app *Application) delete_project (w http.ResponseWriter, r *http.Request){
    id, err := app.readIDparam(r)
    if err != nil {
        http.NotFound(w, r)
    }
    fmt.Fprintf(w, "deleted project of id %d \n", id)
}
