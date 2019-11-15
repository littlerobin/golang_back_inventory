package main

// By go-micro code generator .
// Update by Emran Hamdan  30/9/2019
// main.go is the main service starting point , each micro service define must have
// which  will address gPRC using inventory "inventory-service/proto/inventory" bellow
// We added datastore as part of our design.
import (
	"fmt"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"

	//"github.com/micro/go-plugins/wrapper/monitoring/prometheus"
	"inventory-service/handler"
	//"inventory/subscriber"

	"inventory-service/db/mgodb"
	inventory "inventory-service/proto/inventory"
)

const (
	port        = ":27017"
	defaultHost = "mongodb://localhost"
	database    = "inventory-services"
)

func main() {
	// New Service we will load from config.toml later

	service := micro.NewService(
		micro.Name("go.micro.srv.inventory"),
		//micro.WrapHandler(prometheus.NewHandlerWrapper()),
	)

	// Initialise service
	service.Init()

	// New datastore  for mongoddb URI

	uri := `mongodb://10.0.75.1:27017`

	// initialise datastore to work with our handler

	//dao, err := datastore.NewDataService(uri, database)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// create new handelr for inventory  using datastore sevices
	// init database with layer
	mgodb.NewDB(uri, database)
	// Emran: change from go micro to use our new handler with datastore
	// Register Handler
	//inventory.RegisterInventoryHandler(service.Server(), new(handler.inventory))

	//.RegisterInventoryHandler(service.Server(), inventoryHandler)

	err := inventory.RegisterInventoryHandler(service.Server(), handler.New())
	fmt.Println("ERROR:", err)
	// This is not required and work for other business cases
	// Register Struct as Subscriber
	// micro.RegisterSubscriber("go.micro.srv.inventory", service.Server(), new(subscriber.inventory))

	// Register Function as Subscriber
	// micro.RegisterSubscriber("go.micro.srv.inventory", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
