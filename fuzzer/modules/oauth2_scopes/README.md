## API Module OAuth2 Scopes
This modules fuzzes OpenAPI endpoints, and ensures that OAuth2 scopes are implemented correctly.
Given a list of client ids and secrets with their associated scopes, this module will request the provided 
authorization server for an access token using client credentials flow, then check each endpoint.
There are two criterias for failure:

1. An endpoint does not grant access to a user with the required scopes
2. An endpoint grants access to a user with some or none of the required scopes
