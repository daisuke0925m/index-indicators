package main

import (
	"fmt"
	"index-indicator-apis/server/app/controllers"
	"index-indicator-apis/server/app/models"
	"index-indicator-apis/server/config"
	"index-indicator-apis/server/db"
	"log"
	"net/http"
)

func run() error {
	models.AutoMigrate()
	db.InitRedis()
	go controllers.StreamIngestionData() //TODO

	models, err := models.NewModels()
	if err != nil {
		return err
	}
	defer models.DB.Close()

	app := controllers.NewApp(models)
	r := controllers.Route(app)
	http.Handle("/", r)
	fmt.Printf("connected port :%d|\n", config.Config.Port)
	return http.ListenAndServe(fmt.Sprintf(":%d", config.Config.Port), nil)
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
