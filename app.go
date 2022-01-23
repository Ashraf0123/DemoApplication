package main

import (
	"database/sql"
	"fmt"
	"os"
	v1 "tempApi/api/v1"
	"tempApi/config"
	"tempApi/router"
	"tempApi/server"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	godotenv.Load(".env")
	psqlconn := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable", os.Getenv("host"), os.Getenv("port"), os.Getenv("user"), os.Getenv("password"), os.Getenv("dbname"))
	v1.Db, v1.Err = sql.Open("postgres", psqlconn)

	//init envVars
	//init db connections
	//init certificates

	srv := server.New()
	router.Routes(srv)

	srv.Logger.Fatal(srv.Start(":" + config.ServerPort))
}
