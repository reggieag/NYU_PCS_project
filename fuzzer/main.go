package main

import (
	"context"
	"fmt"
	"fuzzer/pkg/runner"
	"fuzzer/utilities"
	"io/ioutil"
	"log"
	"net/url"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		log.Fatalf("No config file found. Exiting\n")
	}
	fileName := args[1]
	config, err := parseConfig(fileName)
	if err != nil {
		log.Fatalf("Unable to read config: %s", err)
	}

	apiSchema, err := readSchema(config.Runner.API.Schema)
	if err != nil {
		log.Fatalf("Unable to read schema: %s", err)
	}
	apiURL := generateAPIURL(config)
	log.Printf("Using API URL: %s", apiURL)

	log.Printf("Reading clients file")
	clients, err := readClientFile(config.Runner.API.Security.ClientsFile)
	if err != nil {
		log.Fatalf("Unable to read clients: %s", err)
	}

	cliRunner := runner.NewRunner(apiSchema, apiURL, clients, config.ModulesParsed)

	ctx := context.Background()

	startFunc := createStart(config.Runner.ControlScript)
	stopFunc := createStop(config.Runner.ControlScript)

	moduleResults := cliRunner.Execute(ctx, startFunc, stopFunc)

	var hasErrored bool
	for _, result := range moduleResults {
		if result.Error != nil {
			log.Printf("Module %s received error: %s", result.Module.Name, result.Error)
			hasErrored = true
		} else {
			log.Printf("Module %s ran succesfully", result.Module.Name)
		}
	}
	if hasErrored {
		os.Exit(1)
	}
}

func readSchema(file string) (string, error) {
	schema, err := ioutil.ReadFile(file)
	return string(schema), err
}

func generateAPIURL(config *Config) string {
	url := url.URL{
		Scheme: config.Runner.API.HTTPScheme,
		Host:   fmt.Sprintf("%s:%d", config.Runner.API.Host, config.Runner.API.Port),
	}
	return url.String()
}

func createStart(control string) runner.StartSetup {
	return func(ctx context.Context, module runner.ModuleConfig) error {
		log.Printf("Running module %s", module.Name)
		log.Println("Calling control script with argument 'start' for run:")
		if err := utilities.StartAPI(control); err != nil {
			log.Printf("unable to run control script to start API: %s", err)
			return err
		}
		return nil
	}
}

func createStop(control string) runner.StopTeardown {
	return func(ctx context.Context, module runner.ModuleConfig, runError error) {
		if runError != nil {
			log.Printf("Module run complete with error: %s", runError)
		} else {
			log.Printf("Module run complete")
		}
		log.Println("Calling control script with argument 'stop'")
		err := utilities.StopAPI(control)
		if err != nil {
			log.Println(err)
		}
	}
}
