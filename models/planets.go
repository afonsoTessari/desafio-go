package models

type Planet struct {
	Id      interface{} `json:"id,omitempty" bson:"_id,omitempty"`
	Name    string      `json:"name" bson:"name"`
	Climate string      `json:"climate" bson:"climate"`
	Terrain string      `json:"terrain" bson:"terrain"`
	Movies  int         `json:"movies" bson:"movies"`
}

type PlanetDelete struct {
	DeletedCount int64 `json:"deletedCount"`
}
