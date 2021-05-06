package main

import (
	"context"
	"net/http"
	"starwars/config"
	"starwars/database"
	"starwars/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	conf := config.GetConfig()
	ctx := context.TODO()

	db := database.ConnectDB(ctx, conf.Mongo)
	collection := db.Collection(conf.Mongo.Collection)
	client := &database.PlanetClient{
		Collection: collection,
		Ctx:        ctx,
	}

	router := gin.Default()
	planets := router.Group("/planets")

	planets.GET("/", handlers.SearchPlanet(client))
	planets.GET("/{id}", handlers.GetPlanet(client))
	planets.POST("/", handlers.InsertPlanet(client))
	planets.DELETE("/{id}", handlers.DeletePlanet(client))

	http.ListenAndServe(":8080", router)
}
