package openapi

import "context"

type contextKey string

const scopeKey contextKey = "scopes"

func withScopes(ctx context.Context, scopes []string) context.Context {
	if scopes == nil {
		scopes = []string{}
	}
	return context.WithValue(ctx, scopeKey, scopes)
}

func getScopes(ctx context.Context) []string {
	scopesInterface := ctx.Value(scopeKey)
	switch scopesInterface := scopesInterface.(type) {
	case []string:
		return scopesInterface
	default:
		return []string{}
	}
}
