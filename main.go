package main

import (
	"Chess/server/infrastructure"
	"Chess/server/router"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

// @title Swagger Chess project API
// @version 2.0
// @description This is list api for chess project

// @host chess-apis.herokuapp.com
// @BasePath /api/v1/be

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	port := os.Getenv("PORT")
	if port == ""{
		port = "4000"
	}
	fmt.Println("server :", port)

	//to create table in db
	//go run main.go -db=init
	var value string
	flag.StringVar(&value, "db", "", "init")
	flag.Parse()
	if value == "init" {
		db, _ := infrastructure.OpenConnection()
		if err := infrastructure.InitDatabase(); err != nil {
			log.Println("connection problem ", err)
		}
		infrastructure.CloseConnection(db)
	}

	log.Fatal(http.ListenAndServe(":" + port, router.Router()))
}
