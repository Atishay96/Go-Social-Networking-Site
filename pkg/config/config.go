package config

import (
	"fmt"
	"log"
	"os"

	"Go-Social/pkg"

	"github.com/joho/godotenv"
)

func GetConfig() *root.Config {
	return &root.Config{
		Mongo: &root.MongoConfig{
			// store in enviorment instead of here
			Ip:     envOrDefaultString("mongoLink", "127.0.0.1:27017"),
			DbName: envOrDefaultString("dbName", "myDb")},
		Server: &root.ServerConfig{Port: envOrDefaultString("port", ":1377")},
		Auth:   &root.AuthConfig{Secret: envOrDefaultString("secret", "mysecret")}}
}

func envOrDefaultString(envVar string, defaultValue string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	env := os.Getenv("go_env")
	if env == "" {
		env = os.Getenv("default_env")
		fmt.Println(os.Getenv("default_env"))
		fmt.Println(env)
	}
	fmt.Println(env+"_"+envVar, len(env))
	value := os.Getenv(env + "_" + envVar)
	if value == "" {
		return defaultValue
	}

	return value
}
