package services

import (
	"cobadulu/models"
	"cobadulu/utils"
	"database/sql"
	"fmt"
	"log"
)

var (
	dbUser, dbPass, dbHost, dbPort, dbSchema string
)

func InitDB() *sql.DB {
	dbUser = utils.GetCustomConf("DbUser", "root")
	dbPass = utils.GetCustomConf("DbPass", "password")
	dbHost = utils.GetCustomConf("DbHost", "localhost")
	dbPort = utils.GetCustomConf("DbPort", "3306")
	dbSchema = utils.GetCustomConf("DbSchema", "schema")

	dbPath := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbSchema)
	db, err := models.ConnectDB(dbPath)
	if err != nil {
		log.Panic(err)
	}
	return db
}
