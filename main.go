package main

import (
	"api"
	"db"
	"log"
)

func main() {
	pgStorageParams := db.NewStorageParams("postgres", "postgres", "5432", "postgres", "postgres", "disable")
	pgStorage, err := db.CreateDBConnection(pgStorageParams)
	if err != nil {
		log.Fatalf("failed to create db connection %s", err)
	}
	apiServer := api.MakeApiServer(":8686", *pgStorage)
	apiServer.HandleEndpoints()
}
