package configs

import (
	"database/sql"
	"log"

	// mysql
	_ "github.com/go-sql-driver/mysql"
)

func DatabaseInitialize(configs *EnvConfig) *sql.DB {
	db, err := sql.Open(configs.DBDriver, configs.DBDsn)
	if err != nil {
		panic(err)
	}
	log.Printf("data base %s loaded", configs.DBName)
	return db
}
