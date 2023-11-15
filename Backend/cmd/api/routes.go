package main

import (

	"github.com/go-chi/chi/v5"
)

func (app *Application) routes () *chi.Mux{
    r := chi.NewRouter()
    r.Route("/v1", func(r chi.Router) {
        r.Route("/projects",func(r chi.Router) {
            r.Get("/", app.all_projects)
            r.Get("/{id}", app.single_project)
            r.Get("/pm/{id}", app.pm_projects)
            r.Get("/assignee/{id}", app.employee_projects)
            r.Get("/employee/{id}/completed", app.employee_completed_projects)

            r.Put("/{id}/completed", app.mark_complete)
            r.Put("/{id}/incomplete", app.mark_uncomplete)
        })
    })
    
    return r
}
