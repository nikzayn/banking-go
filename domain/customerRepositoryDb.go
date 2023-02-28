package domain

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type CustomerRepositoryDb struct {
}

func (d CustomerRepositoryDb) FindAll() ([]Customer, error) {
	db, err := sql.Open("mysql", "root:nikhil123")
}
