package main

import (
	"api"
	"db"
)

func main() {
	apiServer := api.MakeApiServer(":8686", db.Storage{})
	apiServer.HandleEndpoints()
}
