package db

import (
	"log"

	"github.com/UltimateForm/gopen-api/internal/config"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

var Driver neo4j.DriverWithContext

func init() {
	dbCfg := config.LoadDbConfig()
	driver, driverErr := neo4j.NewDriverWithContext(dbCfg.Uri, neo4j.BasicAuth(dbCfg.Username, dbCfg.Password, ""))
	if driverErr != nil {
		log.Println(driverErr)
	}
	Driver = driver
	log.Println("DB PACKAGE INIT")
}
