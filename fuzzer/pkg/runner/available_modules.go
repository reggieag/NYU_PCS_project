package runner

import (
	"fmt"
	"fuzzer/modules/oauth2_scopes"
	"fuzzer/modules/sql_injection"
)

var availableModules = map[string]moduleFunc{
	"sql_injection": sql_injection.SQLInjectorModule,
	"oauth2_scopes": oauth2_scopes.OAuth2ScopesModule,
}

var ModuleNotFound = fmt.Errorf("module does not exist")
