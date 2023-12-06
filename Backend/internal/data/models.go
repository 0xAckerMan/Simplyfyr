package data

import (
	"database/sql"
	"errors"
)

var(
    ErrRecordNotFound = errors.New("record not found")
)

type Models struct{
    Projects ProjectModel
    Roles RoleModel
}

func NewModel(db *sql.DB) Models{
    return Models{
        Projects: ProjectModel{DB: db},
        Roles: RoleModel{DB: db},
    }
}
