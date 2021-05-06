package models

type Planet struct {
	Name    string      `json:"Name" bson:"Name"`
	ID      interface{} `json:"id,omitempty" bson:"_id,omitempty`
	Climate string      `json:"Climate" bson:"Climate"`
	Terrain string      `json:"Terrain" bson:"Terrain"`
	Movies  int         `json:"Movies" bson:"Movies"`
}

type PlanetDelete struct {
	DeletedCount int64 `json:"deletedCount"`
}
