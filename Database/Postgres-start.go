package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"strconv"
	"sync"
	"tempApi/config"

	"github.com/joho/godotenv"
)

var master *sql.DB
var slave [50]*sql.DB
var count = 0
var mtx sync.Mutex
var error2 error

func InitPostgresMaster() error {
	godotenv.Load(".env")
	psqlconn := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable", os.Getenv("host"), os.Getenv("port"), os.Getenv("user"), os.Getenv("password"), os.Getenv("dbname"))
	master, error2 = sql.Open("postgres", psqlconn)
	if error2 != nil {
		return errors.New("error connecting master")
	}
	fmt.Println("master instance loaded")
	return nil
}
func InitPostgresSlave() error {
	for i := 0; i < config.PostgresSlaveCount; i++ {
		psqlconn := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable", os.Getenv("host"), os.Getenv("port"), os.Getenv("user"), os.Getenv("password"), os.Getenv("dbname"))
		slave[i], error2 = sql.Open("postgres", psqlconn)
		if error2 != nil {
			return errors.New("error connecting slave: " + strconv.Itoa(i))
		}
	}
	fmt.Println("slave instance loaded: " + strconv.Itoa(config.PostgresSlaveCount))
	return nil
}

func GetPostgresMaster() *sql.DB {
	return master
}
func GetPostgresSlave() *sql.DB {
	if config.PostgresSlaveCount == 0 {
		return master
	}
	instance := slave[count]
	mtx.Lock()
	count++
	count = count % config.PostgresSlaveCount
	mtx.Unlock()
	return instance
}
