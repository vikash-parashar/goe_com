package db

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/yourusername/ecommerce/config"
)

var DB *sql.DB

func InitDB() error {
	var err error
	DB, err = sql.Open("postgres", config.AppConf.DBConnectionString)
	if err != nil {
		return err
	}

	err = DB.Ping()
	if err != nil {
		return err
	}

	return nil
}
