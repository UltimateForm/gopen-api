package db

import (
	"log"

	"github.com/UltimateForm/gopen-api/internal/config"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func init() {
	log.Println("DB PACKAGE INIT")
}

func GetDriver() (neo4j.DriverWithContext, error) {
	dbCfg := config.LoadDbConfig()
	driver, driverErr := neo4j.NewDriverWithContext(dbCfg.Uri, neo4j.BasicAuth(dbCfg.Username, dbCfg.Password, ""))
	return driver, driverErr
}
