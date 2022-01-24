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
	//godotenv.Load(".env")
	//psqlconn := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable", os.Getenv("host"), os.Getenv("port"), os.Getenv("user"), os.Getenv("password"), os.Getenv("dbname"))
	//v1.Db, v1.Err = sql.Open("postgres", psqlconn)
	err := config.InitEnvironmentVariables()
	if err != nil {
		log.Fatal("envVars error: " + err.Error())
	}
	err = postgres.InitPostgresMaster()
	if err != nil {
		log.Fatal("master endpoint error: " + err.Error())
	}
	err = postgres.InitPostgresMaster()
	if err != nil {
		log.Fatal("slave endpoint error: " + err.Error())
	}

	//init envVars
	//init db connections
	//init certificates

	srv := server.New()
	router.Routes(srv)

	srv.Logger.Fatal(srv.Start(":" + config.ServerPort))
}
