package data

import (
	"database/sql"
	"errors"
)

var(
    ErrRecordNotFound = errors.New("record not found")
    ErrEditConflict= errors.New("edit conflict")
)

type Models struct{
    Projects ProjectModel
    Roles RoleModel
    Users UserModel
}

func NewModel(db *sql.DB) Models{
    return Models{
        Projects: ProjectModel{DB: db},
        Roles: RoleModel{DB: db},
        Users: UserModel{DB: db},
    }
}
