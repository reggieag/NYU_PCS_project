package oauth2_scopes

import (
	"context"
	"fuzzer/utilities"
	"log"
	"strconv"

	"gopkg.in/yaml.v2"
)

type oauth2ScopesConfig struct {
	Exhaustive bool `yaml:"exhaustive"`
}

const dockerName = "fuzzer-modules-oauth2_scopes"

func OAuth2ScopesModule(ctx context.Context, moduleConfig interface{}, apiUrl, apiSchema, clients string) error {
	data, _ := yaml.Marshal(moduleConfig)
	config := &oauth2ScopesConfig{}
	yaml.Unmarshal(data, config)
	env := map[string]string{
		"API_URL":     apiUrl,
		"API_SCHEMA":  apiSchema,
		"API_CLIENTS": clients,
		"EXHAUSTIVE":  strconv.FormatBool(config.Exhaustive),
	}
	log.Printf("Starting OAuth2 Scopes Module Container\n")
	return utilities.RunImage(dockerName, env)

}
