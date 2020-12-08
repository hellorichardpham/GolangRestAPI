package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func CreateDatabase() (*sql.DB, error) {
	// serverName := "localhost:3306"
	// user := "myuser"
	// password := "pw"
	// dbName := "demo"

	//connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=true&multiStatements=true", user, password, serverName, dbName)
	connectionString := "myuser:thisismypassword@tcp(localhost:3306)/demo"

	db, err := sql.Open("mysql", connectionString)

	if err != nil {
		return nil, err
	}

	return db, nil
}
