package main

import (
	"net/http"
	"os"
    "log"
	"github.com/urfave/negroni"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func specificHandler(w http.ResponseWriter, r *http.Request) {

}

func handleEnv() {
	err := godotenv.Load()
  	if err != nil {
    	log.Fatal("Error loading .env file")
  	}
}

func init() {
	// get .env files
	handleEnv()
}

func main() {

	// mux handle routes
	router := mux.NewRouter()
	router.HandleFunc("/specific", specificHandler)

	// config
	env := os.Getenv("GO_ENV")
	if "" == env {
	  env = "Development"
	}

	// appending middlewares
	server := negroni.Classic()

	// router handler with negroni
	server.UseHandler(router)

	// starting server
	server.Run(":" + os.Getenv(env + "_PORT"))
}