package config

import "os"

type APIConfigs struct {
	Key string
}

func GetAPIKey() *APIConfigs {
	return &APIConfigs{
		Key: os.Getenv("X_RAPIDAPI_KEY"),
	}
}
