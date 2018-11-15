package model

import (
	"database/sql"
	"fmt"
)

type Costumer struct {
	IdCostumer int64  `json:"id_costumer"`
	Name       string `json:"name"`
	Address    string `json:"address"`
	Email      string `json:"email"`
	UserName   string `json:"username"`
	Password   string `json:"password"`
}

func (c *Costumer) AddCostumer(db *sql.DB) (int64, error) {
	var id int64
	query := `INSERT INTO ecommerce.costumer (name,address,email,username,password) VALUES ($1,$2,$3,$4,$5) RETURNING id_costumer`
	err := db.QueryRow(fmt.Sprintf(query), c.Name, c.Address, c.Email, c.UserName, c.Password).Scan(&id)
	return id, err
}

func (c *Costumer) GetAllCostumer(db *sql.DB) ([]Costumer, error) {
	costumers := []Costumer{}
	query := `SELECT id_costumer,name,address,email,username,password FROM ecommerce.costumer`
	rows, err := db.Query(fmt.Sprintf(query))
	if err != nil {
		return costumers, err
	}

	for rows.Next() {
		costumer := Costumer{}
		rows.Scan(&costumer.IdCostumer,
			&costumer.Name,
			&costumer.Address,
			&costumer.Email,
			&costumer.UserName,
			&costumer.Password,
		)
		costumers = append(costumers, costumer)
	}

	return costumers, nil
}

func (c *Costumer) GetOneCostumer(db *sql.DB) (*Costumer, error) {
	costumer := &Costumer{}
	query := `SELECT id_costumer,name,address,email,username,password FROM ecommerce.costumer WHERE id_costumer=$1`
	err := db.QueryRow(fmt.Sprintf(query), c.IdCostumer).Scan(&costumer.IdCostumer,
		&costumer.Name,
		&costumer.Address,
		&costumer.Email,
		&costumer.UserName,
		&costumer.Password,
	)
	return costumer, err
}

func (c *Costumer) GetLoginCostumer(db *sql.DB) (*Costumer, error) {
	costumer := &Costumer{}
	query := `SELECT id_costumer,name,address,email,username,password FROM ecommerce.costumer WHERE username=$1 OR email=$2`
	err := db.QueryRow(fmt.Sprintf(query), c.UserName, c.Email).Scan(&costumer.IdCostumer,
		&costumer.Name,
		&costumer.Address,
		&costumer.Email,
		&costumer.UserName,
		&costumer.Password,
	)
	if err != nil {
		return costumer, err
	}
	if c.Password == costumer.Password {
		return costumer, nil
	}
	return &Costumer{}, nil
}
