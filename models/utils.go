package models

type Movie struct {
	Result []Result `json:"results"`
}

type Result struct {
	Name   string   `json:"name"`
	Movies []string `json:"films"`
}
