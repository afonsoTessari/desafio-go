package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"starwars/database"
	"starwars/handlers"
	"starwars/models"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestSearchPlanetById(t *testing.T) {
	client := &database.MockPlanetClient{}
	id := primitive.NewObjectID().Hex()
	tests := map[string]struct {
		id           string
		expectedCode int
	}{
		"should return 200": {
			id:           id,
			expectedCode: 200,
		},
		"should return 404": {
			id:           "",
			expectedCode: 404,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			if test.expectedCode == 200 {
				client.On("SearchById", test.id).Return(models.Planet{}, nil)
			}

			req, _ := http.NewRequest("GET", "/planets/"+test.id, nil)
			rec := httptest.NewRecorder()

			r := gin.Default()
			r.GET("/planets/:id", handlers.SearchPlanetById(client))
			r.ServeHTTP(rec, req)

			if test.expectedCode == 200 {
				client.AssertExpectations(t)
			} else {
				client.AssertNotCalled(t, "Get")
			}
		})
	}
}

func TestSearchPlanet(t *testing.T) {
	client := &database.MockPlanetClient{}
	tests := map[string]struct {
		payload      string
		expectedCode int
		expected     string
	}{
		"should return 200 - found": {
			payload:      `{"name":"Kalee"}`,
			expectedCode: 200,
			expected:     "Kalee",
		},
		"should return 400": {
			payload:      "invalid json string",
			expectedCode: 400,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {

			client.On("Search", mock.Anything).Return([]models.Planet{}, nil)

			req, _ := http.NewRequest("GET", "/planets?name="+test.payload, nil)
			rec := httptest.NewRecorder()

			r := gin.Default()
			r.GET("/planets", handlers.SearchPlanet(client))
			r.ServeHTTP(rec, req)

			if test.expectedCode == 200 {
				client.AssertExpectations(t)
			} else {
				client.AssertNotCalled(t, "Search")
			}
		})
	}
}
