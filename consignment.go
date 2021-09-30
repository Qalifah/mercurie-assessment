package mercurie_assessment

import "context"

// Item represent bought goods
type Item struct {
	Name 	string		`json:"name"`
	Price	uint64		`json:"price"`
}

// Consignment represent shipped items
type Consignment struct {
	ID 			 	string		`json:"id"`
	Name			string		`json:"name"`
	Description  	string		`json:"description"`
	Items			[]*Item		`json:"items"`
}

// Repository represents the methods the database needs to implement
type Repository interface {
	Create(ctx context.Context, consignment *Consignment) (string, error)
	Get(ctx context.Context, id string) (*Consignment, error)
	GetAll(ctx context.Context) ([]*Consignment, error)
	Update(ctx context.Context, id string, changes map[string]interface{}) error
	Delete(ctx context.Context, id string) error
}