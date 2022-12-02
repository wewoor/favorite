package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"

	"favorite/ent"

	_ "github.com/go-sql-driver/mysql"
)

type DBClient struct{}

var dbClientInstance *ent.Client
var once sync.Once

func GetClientDms() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
}

func GetDBClient() *ent.Client {
	if dbClientInstance == nil {
		once.Do(func() {
			dbDms := GetClientDms()

			log.Printf("Get Database Client via DMS:" + dbDms)

			dbClient, err := ent.Open("mysql", dbDms)
			dbClientInstance = dbClient
			if err != nil {
				log.Fatalf("failed opening connection to mysql: %v", err)
			}
			// defer dbClient.Close()
			// Run the auto migration tool.
			if err := dbClient.Schema.Create(context.Background()); err != nil {
				log.Fatalf("failed creating schema resources: %v", err)
			}
		})
		return dbClientInstance
	}
	return dbClientInstance
}

func CloseDBClient() {
	if dbClientInstance != nil {
		dbClientInstance.Close()
	}
}
