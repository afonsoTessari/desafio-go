package database

import (
	"starwars/models"

	"github.com/stretchr/testify/mock"
)

type MockPlanetClient struct {
	mock.Mock
}

func (m *MockPlanetClient) Insert(planet models.Planet) (models.Planet, error) {
	args := m.Called(planet)
	return args.Get(0).(models.Planet), args.Error(1)
}

func (m *MockPlanetClient) Delete(id string) (models.PlanetDelete, error) {
	args := m.Called(id)
	return args.Get(0).(models.PlanetDelete), args.Error(1)
}

func (m *MockPlanetClient) Search(filter interface{}) ([]models.Planet, error) {
	args := m.Called(filter)
	return args.Get(0).([]models.Planet), args.Error(1)
}

func (m *MockPlanetClient) SearchById(id string) (models.Planet, error) {
	args := m.Called(id)
	return args.Get(0).(models.Planet), args.Error(1)
}
