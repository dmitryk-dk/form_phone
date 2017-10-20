package database

import (
	"database/sql"
	"fmt"

	"github.com/dmitryk-dk/from_phone/server/config"
	"github.com/dmitryk-dk/from_phone/server/models"
)

var dbInstance *sql.DB

func Connect(config *config.Config) (*sql.DB, error) {
	if dbInstance == nil {
		connectionConfig := fmt.Sprintf("%s:%s@%s/%s", config.User, config.Password, config.Host, config.DbName)
		db, err := sql.Open(config.DbDriverName, connectionConfig)
		if err != nil {
			return nil, err
		}
		dbInstance = db
	}
	return dbInstance, nil
}

type DbMethodsHelper interface {
	AddPhone(*models.Phone) error
	DeletePhone(*models.Phone) error
	GetPhones() (models.Phones, error)
}

type DbMethods struct{}

func (m *DbMethods) AddPhone(phone *models.Phone) error {
	stmt, err := dbInstance.Prepare("INSERT blacklist SET msisdn=?")
	if err != nil {
		fmt.Errorf("Can't add to database: %s", err)
		return nil
	}
	res, err := stmt.Exec(phone.Number)
	if err != nil {
		fmt.Errorf("Can't add to database: %s", err)
		return nil
	}
	fmt.Printf("%v\n", res)
	return nil
}

func (m *DbMethods) DeletePhone(phone *models.Phone) error {
	stmt, err := dbInstance.Prepare("DELETE from blacklist where msisdn=?")
	if err != nil {
		fmt.Errorf("Can't delete from database: %s", err)
		return nil
	}
	res, err := stmt.Exec(&phone.Number)
	if err != nil {
		fmt.Errorf("Can't add to database: %s", err)
		return nil
	}
	fmt.Printf("%v\n", res)
	return nil
}

func (m *DbMethods) GetPhones() (models.Phones, error) {
	phones := make(models.Phones, 0)
	rows, err := dbInstance.Query("SELECT * FROM blacklist")
	if err != nil {
		fmt.Errorf("Can't add to database: %s", err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		phone := &models.Phone{}
		rows.Scan(&phone.Number)
		phones = append(phones, *phone)
	}
	return phones, nil
}
