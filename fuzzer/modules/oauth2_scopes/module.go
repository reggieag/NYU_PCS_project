package oauth2_scopes

import (
	"context"
	"fuzzer/utilities"
	"log"
)

const dockerName = "fuzzer-modules-oauth2_scopes"

func OAuth2ScopesModule(ctx context.Context, moduleConfig interface{}, apiUrl, apiSchema, clients string) error {
	env := map[string]string{
		"API_URL":     apiUrl,
		"API_SCHEMA":  apiSchema,
		"API_CLIENTS": clients,
	}
	log.Printf("Starting OAuth2 Scopes Module Container\n")
	return utilities.RunImage(dockerName, env)

}
