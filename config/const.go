//package config

//const ServerPort = "4040"
package config

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// ServerPort and Postgres are embedded into actual code, as they are part of the codebase
//If you run Postgres in a different port, please change the default value
//If you want to use a different port for the application, change the server port. Please also make changes to docker file (expose the same port)
const ServerPort = "4040"
const PostgresPort = "5432"

//envVars
var PostgresPassword string
var PostgresMasterEndpoint string
var PostgresSlaveEndpoints [50]string
var PostgresSlaveCount int

//temp variable
var boolVal bool
var slaveCountTemp string

func InitEnvironmentVariables() error {
	godotenv.Load(".env")

	//DB CREDENTIALS + Cluster Endpoint + Others
	PostgresPassword, boolVal = os.LookupEnv("Postgres_PASSWORD")
	if boolVal == false {
		return errors.New("Postgres_PASSWORD not found in envVars")
	}
	PostgresMasterEndpoint, boolVal = os.LookupEnv("MASTER_ENDPOINT")
	if boolVal == false {
		return errors.New("MASTER_ENDPOINT not found in envVars")
	}
	err := initSlaveEndpoints()
	if err != nil {
		return err
	}
	fmt.Println("environment vars loaded")
	return nil
}
func initSlaveEndpoints() error {
	godotenv.Load(".env")
	slaveCountTemp, boolVal = os.LookupEnv("SLAVE_COUNT")
	if boolVal == true {
		var err error
		PostgresSlaveCount, err = strconv.Atoi(slaveCountTemp)
		if err != nil {
			return err
		}
		if PostgresSlaveCount < 0 || PostgresSlaveCount > 50 {
			return errors.New("invalid slave number: " + slaveCountTemp)
		}
	} else {
		PostgresSlaveCount = 0
		return nil
	}
	for i := 0; i < PostgresSlaveCount; i++ {
		PostgresSlaveEndpoints[i], boolVal = os.LookupEnv("SLAVE_ENDPOINT_" + strconv.Itoa(i))
		if boolVal == false {
			return errors.New("SLAVE_ENDPOINT_" + strconv.Itoa(i) + " not found in envVars")
		}
	}
	return nil
}
