package handlers

import (
	"net/http"
	"starwars/database"
	"starwars/models"
	"starwars/utils"

	"github.com/gin-gonic/gin"
)

func InsertPlanet(db database.PlanetInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		planet := models.Planet{}
		err := c.BindJSON(&planet)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		movies, err := utils.GetMovies(planet.Name)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		planet.Movies = movies

		res, err := db.Insert(planet)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, res)
	}
}
