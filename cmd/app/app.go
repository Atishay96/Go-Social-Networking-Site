package main

import (
	"fmt"
	"log"

	"Go-Social/pkg"
	"Go-Social/pkg/config"
	"Go-Social/pkg/mongo"
	"Go-Social/pkg/server"
)

type App struct {
	server  *server.Server
	session *mongo.Session
	config  *root.Config
}

func (a *App) Initialize() {
	a.config = config.GetConfig()
	var err error
	a.session, err = mongo.NewSession(a.config.Mongo)
	if err != nil {
		log.Fatalln("unable to connect to mongodb")
	}

	u := mongo.NewUserService(a.session.Copy(), a.config.Mongo)
	p := mongo.NewPostService(a.session.Copy(), a.config.Mongo)
	a.server = server.NewServer(u, p, a.config)
}

func (a *App) Run() {
	fmt.Println("Welcome back, Atishay")
	defer a.session.Close()
	a.server.Start()
}
