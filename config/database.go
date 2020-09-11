package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/hansenedrickh/katachi/utils"
)

func InitDB() *sql.DB {
	dbHost := utils.FatalGetString("DATABASE_HOST")
	dbPort := utils.FatalGetString("DATABASE_PORT")
	dbUser := utils.FatalGetString("DATABASE_USER")
	dbPass := utils.FatalGetString("DATABASE_PASS")
	dbName := utils.FatalGetString("DATABASE_NAME")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPass, dbName)

	dbConn, err := sql.Open("postgres", dsn)
	if err != nil {
		panic("cannot open DB connection")
	}

	dbConn.SetMaxIdleConns(0)

	err = dbConn.Ping()
	if err != nil {
		panic("cannot connect to DB")
	}

	return dbConn
}
