package model

import (
	"database/sql"
)

type AdminStruct struct {
	IdAdmin  int64  `json:"id_admin"`
	Name     string `json:"name"`
	Address  string `json:"address"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (a *AdminStruct) Login(db *sql.DB) (*AdminStruct, error) {
	result := &AdminStruct{}
	query := `SELECT id_admin,name,address,email,username,password FROM ecommerce.admin WHERE username=$1 AND password=$2`
	err := db.QueryRow(query, a.Username, a.Password).Scan(&result.IdAdmin, &result.Name, &result.Address, &result.Email, &result.Username, &result.Password)
	return result, err
}
