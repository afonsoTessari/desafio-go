package handlers_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"starwars/config"
	"starwars/database"
	"starwars/handlers"
	"starwars/models"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var client database.PlanetInterface

func init() {
	conf := config.MongoConfiguration{
		Server:     "mongodb://localhost:27017",
		Database:   "Mgo",
		Collection: "PlanetsTest",
	}
	ctx := context.TODO()

	db := database.ConnectDB(ctx, conf)
	collection := db.Collection(conf.Collection)

	client = &database.PlanetClient{
		Collection: collection,
		Ctx:        ctx,
	}
}

func TestInsertPlanet(t *testing.T) {
	tests := map[string]struct {
		payload      string
		expectedCode int
		expected     string
	}{
		"should return 200": {
			payload:      `{"name": "Kalee","climate": "seco","terrain": "arid"}`,
			expectedCode: 200,
			expected:     "Kalee",
		},
		"should return 400": {
			payload:      "invalid string",
			expectedCode: 400,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			req, _ := http.NewRequest("POST", "/planets", strings.NewReader(test.payload))
			rec := httptest.NewRecorder()
			h := http.HandlerFunc(handlers.InsertPlanet(client))
			h.ServeHTTP(rec, req)

			if test.expectedCode == 200 {
				planet := models.Planet{}
				_ = json.Unmarshal([]byte(rec.Body.String()), &planet)
				assert.Equal(t, test.expected, planet.Name)
				assert.NotNil(t, planet.Id)

				_, _ = client.Delete(planet.Id.(string))
			}
			assert.Equal(t, test.expectedCode, rec.Code)
		})
	}
}
