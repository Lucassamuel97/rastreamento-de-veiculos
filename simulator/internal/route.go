package internal

import (
	"fmt"
	"math"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Directions struct {
	Lat float64 `bson:"lat" json:"lat"`
	Lng float64 `bson:"lng" json:"lng"`
}

type Route struct {
	ID           string       `bson:"_id" json:"id"`
	Distance     int          `bson:"distance" json:"distance"`
	Directions   []Directions `bson:"directions" json:"directions"`
	FreightPrice float64      `bson:"freight_price" json:"freight_price"`
}

func NewRoute(id string, distance int, directions []Directions) Route {
	return Route{
		ID:         id,
		Distance:   distance,
		Directions: directions,
	}
}

type RouteService struct {
	mongo          *mongo.Client
	FreightService *FreightService
}

func NewRouteService(mongo *mongo.Client, freightService *FreightService) *RouteService {
	return &RouteService{
		mongo:          mongo,
		FreightService: freightService,
	}
}

type FreightService struct{}

func NewFreightService() *FreightService {
	return &FreightService{}
}

func (fs *FreightService) Calculate(distance int) float64 {
	return math.Floor((float64(distance)*0.15+0.3)*100) / 100
}

func (rs *RouteService) CreateRoute(route Route) (Route, error) {

	// Use FreightService to perform necessary calculations
	freightCost := rs.FreightService.Calculate(route.Distance)
	route.FreightPrice = freightCost
	fmt.Printf("Calculated freight cost: %.2f\n", freightCost)

	update := bson.M{
		"$set": bson.M{
			"distance":      route.Distance,
			"directions":    route.Directions,
			"freight_price": freightCost,
		},
	}

	// Filter to find document by ID
	filter := bson.M{"_id": route.ID}

	// Upsert option to insert if not exists
	opts := options.Update().SetUpsert(true)

	// Perform the update or insert operation
	_, err := rs.mongo.Database("routes").Collection("routes").UpdateOne(nil, filter, update, opts)

	if err != nil {
		return Route{}, err
	}

	return route, err
}

func (rs *RouteService) GetRoute(id string) (Route, error) {
	var route Route
	filter := bson.M{"_id": id}
	err := rs.mongo.Database("routes").Collection("routes").FindOne(nil, filter).Decode(&route)
	fmt.Printf("Found route: %+v\n", route)
	return route, err
}
