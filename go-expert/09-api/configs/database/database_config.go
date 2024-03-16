package database

import (
	"database/sql"
	"github.com/DanielAgostinhoSilva/fullcycle/9-api/configs/enviroment"
	"log"

	// mysql
	_ "github.com/go-sql-driver/mysql"
)

func InitializeDatabase(configs *enviroment.EnvConfig) *sql.DB {
	db, err := sql.Open(configs.DBDriver, configs.DBDsn)
	if err != nil {
		panic(err)
	}
	log.Printf("data base %s loaded", configs.DBName)
	return db
}
