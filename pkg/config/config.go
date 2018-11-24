package config

import (
	"os"

	"Go-Social/pkg"
)

func GetConfig() *root.Config {
	return &root.Config{
		Mongo: &root.MongoConfig{
			// store in enviorment instead of here
			Ip:     envOrDefaultString("Go-Social:mongo:ip", "127.0.0.1:27017"),
			DbName: envOrDefaultString("Go-Social:mongo:dbName", "myDb")},
		Server: &root.ServerConfig{Port: envOrDefaultString("Go-Social:server:port", ":1377")},
		Auth:   &root.AuthConfig{Secret: envOrDefaultString("Go-Social:auth:secret", "mysecret")}}
}

func envOrDefaultString(envVar string, defaultValue string) string {
	value := os.Getenv(envVar)
	if value == "" {
		return defaultValue
	}

	return value
}
