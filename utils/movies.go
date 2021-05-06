package utils

import (
	"encoding/json"
	"net/http"
	"starwars/models"
)

func GetMovies(name string) (int, error) {
	movies := models.Movie{}
	number := 0
	responseSearch, err := http.Get("https://swapi.dev/api/planets/?search=" + name)
	if err != nil {
		return number, err
	}
	defer responseSearch.Body.Close()

	err = json.NewDecoder(responseSearch.Body).Decode(&movies)
	if err != nil {
		return number, err
	}

	if len(movies.Result) > 0 {
		number = len(movies.Result[0].Movies)
	}

	return number, nil
}
