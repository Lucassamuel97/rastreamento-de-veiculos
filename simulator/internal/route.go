package internal

import "go.mongodb.org/mongo-driver/mongo"

type Directions struct {
	Lat float64
	Lng float64
}

type Route struct {
	ID           string
	Distance     int
	Directions   []Directions
	FreightPrice float64
}

type RouteService struct {
	mong *mongo.Client
}
