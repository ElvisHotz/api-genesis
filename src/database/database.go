package database

import (
	"database/sql"
	"teste-api/src/config"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
	db, erro := sql.Open("mysql", config.ConnectionString)
	if erro != nil {
		return nil, erro
	}
	if erro = db.Ping(); erro != nil {
		db.Close()
		return nil, erro
	}
	return db, nil
}
