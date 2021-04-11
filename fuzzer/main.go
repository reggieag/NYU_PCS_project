package main

import (
	"context"
	"fuzzer/modules/sql_injection"
	"fuzzer/utilities"
	"log"
)

const fileName = "../docker-compose.yml"

const apiPort = 8080
const dbPort = 5432
const dbName = "toy_api"
const dbUsername = "api_user"
const dbPassword = "password"

type moduleFunc func(ctx context.Context, apiPort, dbPort int, dbName, dbUsername, dbPassword string) error

var modulesToRun = map[string]moduleFunc{
	"sql_injection": sql_injection.SQLInjectorModule,
}

func main() {
	// Parse configs and stuff
	for name, module := range modulesToRun {
		err := runModule(name, module, fileName)
		if err != nil {
			log.Printf("Error running module %s: %s\n", name, err)
		}
	}
}

func runModule(moduleName string, module moduleFunc, dockerCompose string) error {
	log.Printf("Running module %s\nStarting API/DB for run\n", moduleName)
	utilities.StartAPI(dockerCompose)
	defer func() {
		log.Printf("Module run complete. Shutting down current API/DB\n")
		err := utilities.StopAPI(dockerCompose)
		if err != nil {
			log.Println(err)
		}
	}()
	ctx := context.Background()
	return module(ctx, apiPort, dbPort, dbName, dbUsername, dbPassword)
}
