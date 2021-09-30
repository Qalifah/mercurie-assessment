package handler

import (
	"context"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"

	db "github.com/Qalifah/mercurie-assessment/database"
	pb "github.com/Qalifah/mercurie-assessment/proto/shipping"
)

func TestShippingHandler(t *testing.T) {
	ctx := context.Background()
	client, err := db.NewClient(ctx, "shipping-123")
	if err != nil {
		log.Fatalf("unable to create database client %v", err)
	}
	consignmentCollection := client.Collection("Consignments")
	db := db.NewRepository(consignmentCollection)
	shipping := New(db)
	consignments := []struct {
		Name			string
		Description  	string
		Items		[]*pb.Item
	} {
		{
			"Foodstuffs",
			"", 
			[]*pb.Item {
				{ Name: "Rice", Price: 0},
				{Name: "Beans", Price: 0},
			},
		},
		{
			"Accessories",
			"", 
			[]*pb.Item {
				{ Name: "Watch", Price: 0},
				{Name: "Radio", Price: 0},
			},
		},
	}

	t.Run("Test CreateConsignment Handler", func(t *testing.T) {
		for _, con := range consignments {
			err = shipping.CreateConsignment(ctx, &pb.Consignment{Name: con.Name, Description: con.Description, Items: con.Items}, &pb.Response{})
			assert.NoError(t, err)
		}
	})

	t.Run("Test GetConsignment Handler", func(t *testing.T) {
		err = shipping.GetConsignment(ctx, &pb.SearchParameter{}, &pb.Response{})
		assert.Error(t, err)
	})

	t.Run("Test GatAllConsignments Handler", func(t *testing.T) {
		err = shipping.GetAllConsignments(ctx, &pb.GetAllRequest{}, &pb.Response{})
		assert.NoError(t, err)
	})

	t.Run("Test UpdateConsignment Handler", func(t *testing.T) {
		err = shipping.UpdateConsignment(ctx, &pb.Consignment{}, &pb.Response{})
		assert.Error(t, err)
	})

	t.Run("Test DeleteConsignment Handler", func(t *testing.T) {
		err = shipping.DeleteConsignment(ctx, &pb.SearchParameter{}, &pb.Response{})
		assert.Error(t, err)
	})

	t.Run("Test QuoteConsignment Handler", func(t *testing.T) {
		err = shipping.QuoteConsignment(ctx, &pb.SearchParameter{}, &pb.Response{})
		assert.Error(t, err)
	})
}