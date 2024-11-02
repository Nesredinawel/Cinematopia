package config

import (
	"os"
)

func GetHasuraEndpoint() string {
	return os.Getenv("HASURA_GRAPHQL_ENDPOINT")
}

func GetAdminSecret() string {
	return os.Getenv("HASURA_ADMIN_SECRET")
}
