// Secondary adapter means server side adapter
package domain

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/nikzayn/banking/errs"
	"github.com/nikzayn/banking/logger"
)

// Define CustomerRepositoryDb struct with client as sql db type
type CustomerRepositoryDb struct {
	client *sqlx.DB
}

/*
* Description - FindAll function is to find all the customers from the customers table
* @Output -> (Customer struct, error)
 */
func (db CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {
	var err error
	customers := make([]Customer, 0)
	if status == "" {
		findAllSql := `select * from customers`
		err = db.client.Select(&customers, findAllSql)

	} else {
		findAllSqlByStatus := `select * from customers where status = ?`
		err = db.client.Select(&customers, findAllSqlByStatus, status)
	}

	if err != nil {
		logger.Error("Error while querying customer table " + err.Error())
		return nil, errs.UnexpectedError("Unexpected database error")
	}
	return customers, nil
}

/*
* Description - FindById function to find the customer by id
* @Output -> (Customer struct, error)
 */
func (db CustomerRepositoryDb) ById(id string) (*Customer, *errs.AppError) {
	findById := `select * from customers where customer_id = ?`

	var c Customer
	err := db.client.Get(&c, findById, id)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NotFound("Customer not found")
		} else {
			logger.Error("Error while scanning customer " + err.Error())
			return nil, errs.UnexpectedError("Error customer not found")
		}
	}
	return &c, nil
}

/*
* Description - Creating DB Pool connection
* @Output -> CustomerRepositoryDb struct
 */
func NewCustomerRepositoryDb() CustomerRepositoryDb {
	client, err := sqlx.Open("mysql", "root:nikhil123@tcp(127.0.0.1:3306)/banking")
	if err != nil {
		panic(err)
	}

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return CustomerRepositoryDb{client}
}
