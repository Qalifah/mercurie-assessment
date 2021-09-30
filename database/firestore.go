package database

import  (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"

	core "github.com/Qalifah/mercurie-assessment"
)

func NewClient(ctx context.Context, projectID string) (*firestore.Client, error) {
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return nil, err
	}
	return client, nil
}

type Repository struct {
	collection *firestore.CollectionRef
}

func NewRepository(collection *firestore.CollectionRef) *Repository {
	return &Repository{
		collection: collection,
	}
}

func(c *Repository) Create(ctx context.Context, consignment *core.Consignment) (string, error) {
	doc, _, err := c.collection.Add(ctx, consignment)
	if err != nil {
		log.Fatalf("failed to create new consignment : %v", err)
		return "", err
	}

	return doc.ID, nil
}

func(c *Repository) Get(ctx context.Context, id string) (*core.Consignment, error) {
	docSnap, err := c.collection.Doc(id).Get(ctx)
	if err != nil {
		log.Fatalf("failed to retrieve consignment : %v", err)
		return nil, err
	}

	var consignment *core.Consignment
	err = docSnap.DataTo(consignment)
	if err != nil {
		return nil, err
	}

	return consignment, nil
}

func(c *Repository) GetAll(ctx context.Context) ([]*core.Consignment, error) {
	iter := c.collection.Documents(ctx)
	var consignments []*core.Consignment

	for {
		var consignment *core.Consignment
        collRef, err := iter.Next()
        if err == iterator.Done {
            break
        }
        if err != nil {
            return nil, err
        }
		collRef.DataTo(consignment)
		consignments = append(consignments, consignment)
	}

	return consignments, nil
}

func(c *Repository) Update(ctx context.Context, id string, changes map[string]interface{}) error {
	_, err := c.collection.Doc(id).Set(ctx, changes, firestore.MergeAll)
	if err != nil {
		log.Fatalf("failed to update consignment : %v", err)
		return err
	}
	return nil
}

func(c *Repository) Delete(ctx context.Context, id string) error {
	_, err := c.collection.Doc(id).Delete(ctx)
	if err != nil {
		log.Fatalf("failed to delete consignment : %v", err)
		return err
	}
	return nil
}