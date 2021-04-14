package main

import (
	"context"
	"fuzzer/modules/sql_injection"
	"fuzzer/utilities"
	"log"
	"os"
)

type moduleFunc func(ctx context.Context, moduleConfig interface{}, apiPort, dbPort int, dbName, dbUsername, dbPassword string) error

var availableModules = map[string]moduleFunc{
	"sql_injection": sql_injection.SQLInjectorModule,
}

func main() {
	args := os.Args
	if len(args) < 2 {
		log.Fatalf("No config file found. Existing\n")
	}
	fileName := args[1]
	config, err := ParseConfig(fileName)
	if err != nil {
		log.Fatalf("Unable to read config: %s", err)
	}
	type failedModule struct {
		moduleName string
		err        error
	}
	var failedModules []failedModule
	for name := range config.Modules {
		moduleFunc, ok := availableModules[name]
		if !ok {
			log.Printf("Module not found: %s. Skipping\n", name)
			continue
		}
		err := runModule(name, moduleFunc, config)
		if err != nil {
			failedModules = append(failedModules, failedModule{name, err})
		}
	}
	if len(failedModules) != 0 {
		for i := range failedModules {
			log.Printf("%s failed: %s\n", failedModules[i].moduleName, failedModules[i].err)
		}
		os.Exit(1)
	}
}

func runModule(moduleName string, module moduleFunc, config *Config) error {
	log.Printf("Running module %s\nStarting API/DB for run\n", moduleName)
	utilities.StartAPI(config.Runner.ComposeFile)
	defer func() {
		log.Printf("Module run complete. Shutting down current API/DB\n")
		err := utilities.StopAPI(config.Runner.ComposeFile)
		if err != nil {
			log.Println(err)
		}
	}()
	// This can in theory be used to set global timeout limits
	ctx := context.Background()
	err := module(ctx,
		config.Modules[moduleName],
		config.Runner.API.Port,
		config.Runner.Database.Port,
		config.Runner.Database.Name,
		config.Runner.Database.Username,
		config.Runner.Database.Password)
	if err != nil {
		log.Printf("Received error from runner: %s", err)
	}
	return err
}
