package main

import (
	"database/sql"
	"fmt"
	v1 "tempApi/api/v1"
	"tempApi/config"
	"tempApi/router"
	"tempApi/server"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "123456"
	dbname   = "my_db"
)

func main() {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	v1.Db, v1.Err = sql.Open("postgres", psqlconn)

	//init envVars
	//init db connections
	//init certificates

	srv := server.New()
	router.Routes(srv)

	srv.Logger.Fatal(srv.Start(":" + config.ServerPort))
}
