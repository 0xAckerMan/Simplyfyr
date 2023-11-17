package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/0xAckerMan/Simplifyr/Backend/internal/data"
)

func (app *Application) all_projects(w http.ResponseWriter, r *http.Request) {
	message := "All projects lists"
    err := app.writeJSON(w, http.StatusOK, envelope{"error": message}, nil)
	if err != nil {
        app.serverErrorResponse(w, r, err)
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
		Done:         false,
	}
    err = app.writeJSON(w, http.StatusOK, envelope{"project": project}, nil)
}

func (app *Application) pm_projects(w http.ResponseWriter, r *http.Request) {
    id, err := app.readIDparam(r)
    if err != nil{
        http.NotFound(w,r)
    }
    fmt.Fprintf(w, "all projects belonging to pm of id %d \n", id)
}

func (app *Application) employee_projects (w http.ResponseWriter, r *http.Request){
    id, err := app.readIDparam(r)
    if err != nil{
        app.notFoundResponse(w, r)
    }

    fmt.Fprintf(w, "all employee projects of id %d \n", id)
}

func (app *Application) employee_completed_projects(w http.ResponseWriter, r *http.Request){
    id,err := app.readIDparam(r)
    if err != nil {
        app.notFoundResponse(w, r)
        return
    }
    fmt.Fprintf(w, "Employee of id %d completed projects \n", id)
}

func (app *Application) mark_complete(w http.ResponseWriter, r *http.Request){
    id, err := app.readIDparam(r)
    if err != nil{
        app.notFoundResponse(w, r)
        return
    }
    fmt.Fprintf(w, "project %d marked as completed \n", id)
}

func (app *Application) mark_uncomplete(w http.ResponseWriter, r *http.Request){
    id,err := app.readIDparam(r)
    if err != nil {
        app.notFoundResponse(w, r)
        return
    }
    fmt.Fprintf(w, "project of id %d marked uncomplete \n", id)
}

func (app *Application) create_project(w http.ResponseWriter, r *http.Request){
    var input struct{
        Name string `json:"name"`
        Category []string `json:"category"`
        Excerpt string `json:"excerpt"`
        Description string `json:"description"`
        Assigned_to int `json:"assigned_to"`
        Created_by int `json:"created_by"`
        Due_date time.Time `json:"due_date"`
        Done bool `json:"done"`
    }
    err := app.readJSON(w, r, &input)
    if err != nil {
        app.badRequestResponse(w, r, err)
        return
    }
    err = app.writeJSON(w,200, envelope{"message": input},nil)
    if err != nil {
        app.serverErrorResponse(w, r, err)
        return
    }
}

func (app *Application) update_project (w http.ResponseWriter, r *http.Request){
    id, err := app.readIDparam(r)
    if err != nil{
        app.notFoundResponse(w, r)
        return
    }
    fmt.Fprintf(w, "update project of id %d \n", id)
}

func (app *Application) delete_project (w http.ResponseWriter, r *http.Request){
    id, err := app.readIDparam(r)
    if err != nil {
        app.notFoundResponse(w, r)
        return
    }
    fmt.Fprintf(w, "deleted project of id %d \n", id)
}
