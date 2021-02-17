package main

import (
	"fmt"
	"index-indicators/server/app/controllers"
	"index-indicators/server/app/models"
	"index-indicators/server/db"
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
	fmt.Printf("connected port :%d|\n", 8080)
	return http.ListenAndServe(fmt.Sprintf(":%d", 8080), nil)
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
