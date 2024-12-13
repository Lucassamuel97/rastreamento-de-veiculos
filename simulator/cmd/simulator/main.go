package main

import (
	"context"
	"log"

	"github.com/Lucassamuel97/rastreamento-de-veiculos/simulator/internal"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	mongoStr := "mongodb://admin:admin@localhost:27017/routes?authSource=admin"
	// open new connection with mongodb
	mongoConnection, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoStr))
	if err != nil {
		panic(err)
	}

	/// create a new route service
	freightService := internal.NewFreightService()
	routeService := internal.NewRouteService(mongoConnection, freightService)

	chDriverMoved := make(chan *internal.DriverMovedEvent)
	chFreightCalculated := make(chan *internal.FreightCalculatedEvent)
	kafkaBroker := "localhost:9092"

	freightWriter := &Kafka.Writer{
		Addr:     Kafka.TCP(kafkaBroker),
		Topic:    "freight",
		Balancer: &Kafka.LeastBytes{},
	}

	simulationWriter := &Kafka.Writer{
		Addr:     Kafka.TCP(kafkaBroker),
		Topic:    "simulator",
		Balancer: &Kafka.LeastBytes{},
	}

	routeReader := Kafka.NewReader(Kafka.ReaderConfig{
		Brokers: []string{kafkaBroker},
		Topic:   "route",
		GroupID: "simulator",
	})

	hub := internal.NewEventHub(
		routeService,
		mongoConnection,
		freightWriter,
		simulationWriter,
	)

	for {
		m, err := routeReader.ReadMessage(context.Background())
		if err != nil {
			panic(err)
			continue
		}

		go func(msg []byte) {
			err = hub.HandleEvent(m.Value)
			if err != nil {
				log.Printf("error handling event: %v", err)
			}
		}(m.Value)
	}

}
