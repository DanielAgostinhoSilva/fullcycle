package configs

import (
	"database/sql"
	"github.com/pressly/goose/v3"
	"log"
)

func MigrationUP(configs EnvConfig) {
	gooseDb := getSql(configs)
	defer gooseDb.Close()
	err := goose.Up(gooseDb, configs.MigrationDir)
	if err != nil {
		panic(err.Error())
	}
}

func MigrationDown(configs EnvConfig) {
	gooseDb := getSql(configs)
	defer gooseDb.Close()
	err := goose.DownTo(gooseDb, configs.MigrationDir, 0)
	if err != nil {
		panic(err.Error())
	}
}

func getSql(configs EnvConfig) *sql.DB {
	gooseDB, err := goose.OpenDBWithDriver(configs.DBDriver, configs.DBDsn)
	if err != nil {
		log.Fatal("Erro ao abrir a conexão com o banco de dados:", err)
		panic(err)
	}

	err = goose.SetDialect(configs.DBDriver)
	if err != nil {
		log.Fatal("Erro ao configurar o dialect:", err)
		panic(err)
	}
	return gooseDB
}
