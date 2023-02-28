package domain

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// Define CustomerRepositoryDb struct with client as sql db type
type CustomerRepositoryDb struct {
	client *sql.DB
}

/*
* Description - FindAll function is to find all the customers from the customers table
	* @Output -> (Customer struct, error)
*/
func (db CustomerRepositoryDb) FindAll() ([]Customer, error) {
	findAllSql := `select * from customers`

	//Query rows with findAllSql query
	rows, err := db.client.Query(findAllSql)
	if err != nil {
		log.Println("Error while querying customer table. " + err.Error())
		return nil, err
	}

	customers := make([]Customer, 0)
	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.DateOfBirth, &c.Status, &c.Zipcode)
		if err != nil {
			log.Println("Error while scanning customers. " + err.Error())
			return nil, err
		}
		customers = append(customers, c)
	}
	return customers, nil
}

/*
* Description - Creating DB Pool connection
	* @Output -> CustomerRepositoryDb struct
*/
func NewCustomerRepositoryDb() CustomerRepositoryDb {
	client, err := sql.Open("mysql", "root:nikhil123@tcp(127.0.0.1:3306)/banking")
	if err != nil {
		panic(err)
	}

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return CustomerRepositoryDb{client}
}
