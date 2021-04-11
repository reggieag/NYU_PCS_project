package sql_injection

import (
	"context"
	"fuzzer/utilities"
	"log"
	"strconv"
)

const dockerName = "fuzzer-modules-sql_injection"

func SQLInjectorModule(ctx context.Context, apiPort, dbPort int, dbName, dbUsername, dbPassword string) error {
	env := map[string]string{
		"APIPort":    strconv.Itoa(apiPort),
		"DBPort":     strconv.Itoa(dbPort),
		"DBName":     dbName,
		"DBUsername": dbUsername,
		"DBPassword": dbPassword,
	}
	log.Printf("Starting SQL Injection Module Container\n")
	return utilities.RunImage(dockerName, env)
}
