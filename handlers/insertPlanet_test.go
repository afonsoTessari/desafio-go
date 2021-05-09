package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"starwars/database"
	"starwars/handlers"
	"starwars/models"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

// var client database.PlanetInterface

// func init() {
// 	conf := config.MongoConfiguration{
// 		Server:     "mongodb://localhost:27017",
// 		Database:   "Mgo",
// 		Collection: "PlanetsTest",
// 	}
// 	ctx := context.TODO()

// 	db := database.ConnectDB(ctx, conf)
// 	collection := db.Collection(conf.Collection)

// 	client = &database.PlanetClient{
// 		Collection: collection,
// 		Ctx:        ctx,
// 	}
// }

func TestInsertPlanet(t *testing.T) {
	client := &database.MockPlanetClient{}
	tests := map[string]struct {
		payload      string
		expectedCode int
		expected     string
	}{
		"should return 200": {
			payload:      `{"name":"Alderaan","climate":"seco","terrain":"arid"}`,
			expectedCode: 200,
			expected:     "seco",
		},
		"should return 400": {
			payload:      "invalid string",
			expectedCode: 400,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			client.On("Insert", mock.Anything).Return(models.Planet{}, nil)
			req, _ := http.NewRequest("POST", "/planets/", strings.NewReader(test.payload))
			rec := httptest.NewRecorder()

			r := gin.Default()
			r.POST("/planets/", handlers.InsertPlanet(client))
			r.ServeHTTP(rec, req)

			if test.expectedCode == 200 {
				client.AssertExpectations(t)
			} else {
				client.AssertNotCalled(t, "Insert")
			}
		})
	}
}
