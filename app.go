package main

import (
	"log"
	postgres "tempApi/Database"
	"tempApi/config"
	"tempApi/router"
	"tempApi/server"

	_ "github.com/lib/pq"
)

func main() {

	err := config.InitEnvironmentVariables()
	if err != nil {
		log.Fatal("envVars error: " + err.Error())
	}
	err = postgres.InitPostgresMaster()
	if err != nil {
		log.Fatal("master endpoint error: " + err.Error())
	}
	err = postgres.InitPostgresSlave()
	if err != nil {
		log.Fatal("slave endpoint error: " + err.Error())
	}

	srv := server.New()
	router.Routes(srv)

	srv.Logger.Fatal(srv.Start(":" + config.ServerPort))
}
