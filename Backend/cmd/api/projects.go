package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/0xAckerMan/Simplifyr/Backend/internal/data"
)

func (app *Application) all_projects(w http.ResponseWriter, r *http.Request) {
	projects, err := app.models.Projects.Get_all()
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"projects": projects}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *Application) single_project(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDparam(r)
	if err != nil {
		http.NotFound(w, r)
	}

	project, err := app.models.Projects.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"project": project}, nil)
}

func (app *Application) pm_projects(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDparam(r)
	if err != nil {
		http.NotFound(w, r)
	}
	fmt.Fprintf(w, "all projects belonging to pm of id %d \n", id)
}

func (app *Application) employee_projects(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDparam(r)
	if err != nil {
		app.notFoundResponse(w, r)
	}

	fmt.Fprintf(w, "all employee projects of id %d \n", id)
}

func (app *Application) employee_completed_projects(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDparam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
	fmt.Fprintf(w, "Employee of id %d completed projects \n", id)
}

func (app *Application) mark_complete(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDparam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
	fmt.Fprintf(w, "project %d marked as completed \n", id)
}

func (app *Application) mark_uncomplete(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDparam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
	fmt.Fprintf(w, "project of id %d marked uncomplete \n", id)
}

func (app *Application) create_project(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name        string    `json:"name"`
		Category    []string  `json:"category"`
		Excerpt     string    `json:"excerpt"`
		Description string    `json:"description"`
		Assigned_to int       `json:"assigned_to"`
		Created_by  int       `json:"created_by"`
		Due_date    time.Time `json:"due_date"`
		Done        bool      `json:"done"`
	}
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	project := &data.Project{
		Name:        input.Name,
		Category:    input.Category,
		Excerpt:     input.Excerpt,
		Description: input.Description,
		Assigned_to: input.Assigned_to,
		Created_by:  input.Created_by,
		Due_date:    input.Due_date,
	}

	err = app.models.Projects.Insert(project)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/projects/%d", project.Id))

	err = app.writeJSON(w, http.StatusCreated, envelope{"project": project}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
}

func (app *Application) update_project(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDparam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
	project, err := app.models.Projects.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	var input struct {
		Name        *string    `json:"name"`
		Category    *[]string  `json:"category"`
		Excerpt     *string    `json:"excerpt"`
		Description *string    `json:"description"`
		Assigned_to *int       `json:"assigned_to"`
		Created_by  *int       `json:"created_by"`
		Due_date    *time.Time `json:"due_date"`
		Done        *bool      `json:"done"`
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if input.Name != nil {
		project.Name = *input.Name
	}
	if input.Category != nil {
		project.Category = *input.Category
	}
	if input.Excerpt != nil {
		project.Excerpt = *input.Excerpt
	}
	if input.Description != nil {
		project.Description = *input.Description
	}
    if input.Assigned_to != nil{
        project.Assigned_to = *input.Assigned_to
    }
    if input.Created_by != nil {
        project.Created_by = *input.Created_by
    }
    if input.Due_date != nil {
        project.Due_date = *input.Due_date
    }
    if input.Done != nil {
        project.Done = *input.Done
    }


	err = app.models.Projects.Update(project)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"project": project}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *Application) delete_project(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDparam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
	err = app.models.Projects.Delete(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"message": "Movie successfully deleted"}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
