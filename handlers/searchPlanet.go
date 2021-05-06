package handlers

import (
	"encoding/json"
	"net/http"
	"starwars/database"

	"github.com/gin-gonic/gin"
)

func SearchPlanet(db database.PlanetInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter interface{}
		query := c.Query("name")

		if query != "" {
			err := json.Unmarshal([]byte(query), &filter)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
				return
			}
		}

		res, err := db.Search(filter)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, res)
	}
}
