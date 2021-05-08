package database

import (
	"context"
	"starwars/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PlanetInterface interface {
	Insert(models.Planet) (models.Planet, error)
	SearchById(string) (models.Planet, error)
	Search(interface{}) ([]models.Planet, error)
	Delete(string) (models.PlanetDelete, error)
}

type PlanetClient struct {
	Ctx        context.Context
	Collection *mongo.Collection
}

func (c *PlanetClient) Insert(docs models.Planet) (models.Planet, error) {
	planet := models.Planet{}

	res, err := c.Collection.InsertOne(c.Ctx, docs)
	if err != nil {
		return planet, err
	}
	id := res.InsertedID.(primitive.ObjectID).Hex()
	return c.SearchById(id)
}

func (c *PlanetClient) SearchById(id string) (models.Planet, error) {
	planet := models.Planet{}

	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return planet, err
	}

	err = c.Collection.FindOne(c.Ctx, bson.M{"_id": _id}).Decode(&planet)
	if err != nil {
		return planet, err
	}

	return planet, nil
}

func (c *PlanetClient) Search(filter interface{}) ([]models.Planet, error) {
	planet := []models.Planet{}
	if filter == nil {
		filter = bson.M{}
	}

	cursor, err := c.Collection.Find(c.Ctx, filter)
	if err != nil {
		return planet, err
	}

	for cursor.Next(c.Ctx) {
		row := models.Planet{}
		cursor.Decode(&row)
		planet = append(planet, row)
	}

	return planet, nil
}

func (c *PlanetClient) Delete(id string) (models.PlanetDelete, error) {

	result := models.PlanetDelete{
		DeletedCount: 0,
	}
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return result, err
	}

	res, err := c.Collection.DeleteOne(c.Ctx, bson.M{"_id": _id})
	if err != nil {
		return result, err
	}
	result.DeletedCount = res.DeletedCount
	return result, nil
}
