package data

import (
	"database/sql"
	"errors"
	"time"

)

type Role struct {
	Id        int64     `json:"id"`
	Role      string    `json:"'role"`
	Create_at time.Time `json:"-"`
	Version   int       `json:"version"`
}

type RoleModel struct {
	DB *sql.DB
}

func (r *RoleModel) Insert(role *Role) error {
	query := `INSERT INTO roles (r_role)
    VALUES($1)
    RETURNING r_id, r_created_at, r_version `

	args := []interface{}{role.Role}

	return r.DB.QueryRow(query, args...).Scan(&role.Id, &role.Create_at, &role.Version)
}

func (r *RoleModel) Get(id int64) (*Role, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}
	query := `SELECT r_id, r_role, r_created_at, r_version FROM roles
    WHERE r_id = $1`

	var role Role

	err := r.DB.QueryRow(query, id).Scan(
		&role.Id,
		&role.Role,
		&role.Create_at,
		&role.Version,
	)
	if err != nil {
		switch {
		case errors.Is(err, ErrRecordNotFound):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &role, nil
}

func (r *RoleModel) Get_all() ([]*Role, error){
    query := `SELECT r_id, r_role, r_created_at, r_version
    FROM roles ORDER BY r_id`

    rows, err := r.DB.Query(query)
    if err != nil{
        return nil, err
    }

    var roles []*Role
    for rows.Next(){
        var role Role
        err := rows.Scan(
            &role.Id,
            &role.Role,
            &role.Create_at,
            &role.Version,
            )
        if err != nil {
            return nil, err
        }

        roles = append(roles, &role)
    }
    if err := rows.Err(); err != nil{
        return nil, err
    }
    return roles, nil
}

func (r *RoleModel) Update(role *Role) error {
    query := `UPDATE roles
    SET r_role = $1, r_version = r_version + 1
    WHERE r_id = $2
    RETURNING r_version` 

    args := []interface{}{role.Role, role.Id}

    return r.DB.QueryRow(query, args...).Scan(&role.Version)
}

func (r *RoleModel) Delete(id int64) error {
    if id < 1{
        return ErrRecordNotFound
    }
    query := `DELETE FROM roles
    WHERE r_id = $1`

    result,err := r.DB.Exec(query, id)
    if err != nil{
        return err
    }

    rowsffected, err := result.RowsAffected()
    if err != nil{
        return err
    }

    if rowsffected == 0{
        return ErrRecordNotFound
    }
	return nil
}

