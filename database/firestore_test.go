package database

import (
	"context"
	"testing"
	"log"

	"github.com/stretchr/testify/assert"

	core "github.com/Qalifah/mercurie-assessment"
)

func TestConsignmentDBHandler(t *testing.T) {
	ctx := context.Background()
	client, err := NewClient(ctx, "shipping-123")
	if err != nil {
		log.Fatalf("unable to create database client %v", err)
	}
	consignmentCollection := client.Collection("Consignments")
	db := NewRepository(consignmentCollection)
	consignments := []struct {
		Name			string
		Description  	string
		Items		[]*core.Item
	} {
		{
			"Foodstuffs",
			"", 
			[]*core.Item {
				{ Name: "Rice", Price: 0},
				{Name: "Beans", Price: 0},
			},
		},
		{
			"Accessories",
			"", 
			[]*core.Item {
				{ Name: "Watch", Price: 0},
				{Name: "Radio", Price: 0},
			},
		},
	}

	t.Run("Test Create Handler", func(t *testing.T) {
		for _, con := range consignments {
			_, err := db.Create(ctx, &core.Consignment{
				Name: con.Name,
				Description: con.Description,
				Items: con.Items,
			})
			assert.NoError(t, err)
		}
	})
	
	t.Run("Test Get Handler", func(t *testing.T) {
		_, err := db.Get(ctx, "")
		assert.Error(t, err)
	})

	t.Run("Test GetAll Handler", func(t *testing.T) {
		_, err := db.GetAll(ctx)
		assert.NoError(t, err)
	})

	t.Run("Test Update Handler", func(t *testing.T) {
		err := db.Update(ctx, "", map[string]interface{}{})
		assert.Error(t, err)
	})

	t.Run("Test Delete Handler", func(t *testing.T) {
		err := db.Delete(ctx, "")
		assert.Error(t, err)
	})
}