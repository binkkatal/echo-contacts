package dao

import (
	"database/sql"
	"fmt"
	"os"

	errors "github.com/binkkatal/echo-contacts/errors"

	_ "github.com/go-sql-driver/mysql"
)

const (
	DRIVER      = "mysql"
	DB_USER_KEY = "DB_USER"
	DB_PASSWORD = "DB_PASSWORD"
	DB_NAME_KEY = "ECHO_CONTACTS_DB_NAME"
)

func Connect() (*sql.DB, error) {

	dbUserName := os.Getenv(DB_USER_KEY)

	if dbUserName == "" {
		return nil, errors.DetailedError{
			Err:         errors.ENV_VAR_ERR,
			Description: fmt.Sprintf("key (%s) not set in env", DB_USER_KEY),
		}
	}

	dbName := os.Getenv(DB_NAME_KEY)

	if dbName == "" {
		return nil, errors.DetailedError{
			Err:         errors.ENV_VAR_ERR,
			Description: fmt.Sprintf("key (%s) not set in env", DB_NAME_KEY),
		}
	}

	dbPassword := os.Getenv(DB_PASSWORD)

	dataSource := fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s", dbUserName, dbPassword, dbName)
	fmt.Print(dataSource)
	db, err := sql.Open(DRIVER, dataSource)
	if err != nil {
		return nil, err
	}
	return db, nil
}
