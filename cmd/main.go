package main

import (
	"context"
	"os"

	"github.com/Qalifah/mercurie-assessment/database"
	"github.com/Qalifah/mercurie-assessment/handler"
	pb "github.com/Qalifah/mercurie-assessment/proto/shipping"

	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/logger"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	// create a new service
	service := micro.NewService(
		micro.Name("shipping"),
		micro.Version("latest"),
	)
	
	// initialise flags
	service.Init()

	// setup database
	projectID := os.Getenv("FIRESTORE_PROJECT_ID")
	ctx := context.Background()

	client, err := database.NewClient(ctx, projectID)
	if err != nil {
		logger.Fatalf("Error occured while trying to create a database client : %v", err)
	}
	defer client.Close()
	consignmentCollection := client.Collection("Consignments")

	pb.RegisterShippingServiceHandler(
		service.Server(), handler.New(database.NewRepository(consignmentCollection)),
	)

	// start the service
	if err := service.Run(); err != nil {
		logger.Fatal(err)
	}

}