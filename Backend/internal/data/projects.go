package data

import (
	"database/sql"
	"errors"
	"time"

	"github.com/lib/pq"
)

type Project struct{
    Id int64 `json:"id"`
    Name string `json:"name"`
    Category []string `json:"category"`
    Excerpt string `json:"excerpt"`
    Description string `json:"description"`
    Assigned_to int `json:"assigned_to"`
    Created_by int `json:"created_by"`
    Created_date time.Time `json:"-"`
    Due_date time.Time `json:"Due_date"`
    Done bool `json:"done"`
}

type ProjectModel struct{
    DB *sql.DB
}

func (p ProjectModel) Insert (project *Project) error {
    query := `
        INSERT INTO projects (p_name, p_category,p_excerpt, p_description, p_assigned_to, p_created_by, p_due_date)
        VALUES ($1, $2, $3, $4,$5, $6, $7)
        RETURNING p_id, p_created_date, p_done
    `
    arg := []interface{}{project.Name, pq.Array(project.Category), project.Excerpt, project.Description, project.Assigned_to, project.Created_by, project.Due_date}

    return p.DB.QueryRow(query, arg...).Scan(&project.Id, &project.Created_date, &project.Done)
}

func (p ProjectModel) Get(id int64) (*Project,error){
    if id < 1{
        return nil, ErrRecordNotFound
    }
    query := `
        SELECT p_id, p_name, p_category, p_excerpt, p_description, p_assigned_to,
        p_created_by, p_created_date, p_due_date, p_done FROM projects WHERE p_id=$1
    `

    var project Project

    err := p.DB.QueryRow(query,id).Scan(
        &project.Id,
        &project.Name,
        pq.Array(&project.Category),
        &project.Excerpt,
        &project.Description,
        &project.Assigned_to,
        &project.Created_by,
        &project.Created_date,
        &project.Due_date,
        &project.Done,
        )

    if err != nil{
        switch {
        case errors.Is(err,sql.ErrNoRows):
            return nil, ErrRecordNotFound
        default:
            return nil, err
    }
    }
    return &project, nil
}

func (p *ProjectModel) Get_all() ([]*Project, error) {
    query := `
        SELECT p_id, p_name, p_category, p_excerpt, p_description, p_assigned_to,
        p_created_by, p_created_date, p_due_date, p_done FROM projects ORDER BY p_id
    `

    rows, err := p.DB.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var projects []*Project
    for rows.Next() {
        var project Project
        err := rows.Scan(
            &project.Id,
            &project.Name,
            pq.Array(&project.Category),
            &project.Excerpt,
            &project.Description,
            &project.Assigned_to,
            &project.Created_by,
            &project.Created_date,
            &project.Due_date,
            &project.Done,
        )
        if err != nil {
            return nil, err
        }
        projects = append(projects, &project)
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }

    return projects, nil
}


func (p *ProjectModel) Update(project *Project) error {
    query := ` 
    UPDATE projects
    SET p_name = $1, p_category = $2,p_excerpt = $3, p_description = $4, p_assigned_to = $5, p_created_by = $6, p_due_date = $7, p_done = $8
    WHERE p_id = $9
    RETURNING p_done
    `
    args := []interface{}{
        project.Name,
        pq.Array(project.Category),
        project.Excerpt,
        project.Description,
        project.Assigned_to,
        project.Created_by,
        project.Due_date,
        project.Done,
        &project.Id,
    }
    return p.DB.QueryRow(query,args...).Scan(&project.Done)
}

func (p ProjectModel) Delete (id int64) error {
    if id < 1 {
        return ErrRecordNotFound
    }
    query := `DELETE FROM projects WHERE p_id = $1`

    result, err := p.DB.Exec(query, id)
    if err != nil {
        return err
    }
    
    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return err
    }

    if rowsAffected == 0 {
        return ErrRecordNotFound
    }

    return nil
}
