package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/asim/go-micro/v3/errors"
	"github.com/asim/go-micro/v3/logger"

	core "github.com/Qalifah/mercurie-assessment"
	pb "github.com/Qalifah/mercurie-assessment/proto/shipping"
)

type Shipping struct {
	consignmentRepo core.Repository
}

func New(consignmentRepo core.Repository) *Shipping {
	return &Shipping{
		consignmentRepo: consignmentRepo,
	}
}

func(s *Shipping) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {
	consignment := MarshalConsignment(req)
	consignmentID, err := s.consignmentRepo.Create(ctx, consignment)
	if err != nil {
		logger.Errorf("unable to create consignment: %v", err)
		return errors.FromError(err)
	}

	res.ConsignmentId = consignmentID
	res.Created = true
	return nil
}

func(s *Shipping) GetConsignment(ctx context.Context, req *pb.SearchParameter, res *pb.Response) error {
	consignment, err := s.consignmentRepo.Get(ctx, req.Id)
	if err != nil {
		logger.Errorf("unable to get consignment: %v", err)
		return errors.FromError(err)
	}

	res.Consignment = UnMarshalConsignment(consignment)
	return nil
}

func(s *Shipping) GetAllConsignments(ctx context.Context, req *pb.GetAllRequest, res *pb.Response) error {
	consignments, err := s.consignmentRepo.GetAll(ctx)
	if err != nil {
		logger.Errorf("unable to get all consignments: %v", err)
		return errors.FromError(err)
	}

	res.Consignments = UnMarshalConsignmentCollection(consignments)
	return nil
}

func(s *Shipping) UpdateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {
	consignment := MarshalConsignment(req)
	err := s.consignmentRepo.Update(ctx, req.Id, fromStructToMap(consignment))
	if err != nil {
		logger.Errorf("unable to update consignment: %v", err)
		return errors.FromError(err)
	}

	res.Updated = true
	return nil
}

func(s *Shipping) DeleteConsignment(ctx context.Context, req *pb.SearchParameter, res *pb.Response) error {
	err := s.consignmentRepo.Delete(ctx, req.Id)
	if err != nil {
		logger.Errorf("unable to delete consignment: %v", err)
		return errors.FromError(err)
	}

	res.Deleted = true
	return nil
}

func(s *Shipping) QuoteConsignment(ctx context.Context, req *pb.SearchParameter, res *pb.Response) error {
	consignment, err := s.consignmentRepo.Get(ctx, req.Id)
	if err != nil {
		logger.Errorf("unable to get consignment: %v", err)
		return errors.FromError(err)
	}
	
	totalPrice, err := s.getConsignmentItemsPrice(ctx, consignment)
	if err != nil {
		logger.Errorf("unable to get total consignment items price: %v", err)
		return errors.FromError(err)
	}

	res.TotalPrice = totalPrice
	return nil
}

func (s *Shipping) getConsignmentItemsPrice(ctx context.Context, consignment *core.Consignment) (uint64, error) {
	key := os.Getenv("PRICETREE_API_KEY")
	var totalPrice uint64
	var wg sync.WaitGroup
	errs := core.NewMultiError()

	for _, i := range consignment.Items {
		// Increment the WaitGroup counter
		wg.Add(1)

		go func(item *core.Item) {
			// Decrement the counter when the goroutine completes.
			defer wg.Done()
			itemName := strings.ReplaceAll(item.Name, " ", "+")
			url := fmt.Sprintf(
				"http://www.pricetree.com/dev/api.ashx?storeUri=https://www.amazon.com/s?k=%v&apikey=%v&offset=1", itemName, key,
			)

			request, _ := http.NewRequest("GET", url, nil)
			response, err := http.DefaultClient.Do(request)
			if err != nil {
				logger.Errorf("request to external api failed: %v", err)
				errs.Append(err)
				return
			}

			defer response.Body.Close()
			body, err := ioutil.ReadAll(response.Body)
			if err != nil {
				logger.Errorf("unable to read request body: %v", err)
				errs.Append(err)
				return 
			}

			marshalBody := []map[string]interface{}{}
			err = json.Unmarshal(body, &marshalBody)
			if err != nil {
				logger.Errorf("unable to marshalize request body: %v", err)
				errs.Append(err)
				return
			}
			price := marshalBody[0]["price"].(uint64)
			item.Price = price
			totalPrice += price
		}(i)
	}

	// wait for goroutines to complete
	wg.Wait()

	if errs.HasErrors() {
		return 0, errs
	}

	err := s.consignmentRepo.Update(ctx, consignment.ID, fromStructToMap(consignment))
	if err != nil {
		return 0, err
	}

	return totalPrice, nil
}

func MarshalItem(c *pb.Item) *core.Item {
	return &core.Item{
		Name: c.Name,
		Price: c.Price,
	}
}

func MarshalItemCollection(c []*pb.Item) []*core.Item {
	var items []*core.Item
	for _, item := range c {
		items = append(items, MarshalItem(item))
	}
	return items
}

func UnMarshalItem(c *core.Item) *pb.Item {
	return &pb.Item{
		Name: c.Name,
		Price: c.Price,
	}
}

func UnMarshalItemCollection(c []*core.Item) []*pb.Item {
	var items []*pb.Item
	for _, item := range c {
		items = append(items, UnMarshalItem(item))
	}
	return items
}

func MarshalConsignment(c *pb.Consignment) *core.Consignment {
	return &core.Consignment{
		ID: c.Id,
		Name: c.Name,
		Description: c.Description,
		Items: MarshalItemCollection(c.Items),
	}
}

func UnMarshalConsignment(c *core.Consignment) *pb.Consignment {
	return &pb.Consignment{
		Id: c.ID,
		Name: c.Name,
		Description: c.Description,
		Items: UnMarshalItemCollection(c.Items),
	}
}

func UnMarshalConsignmentCollection(c []*core.Consignment) []*pb.Consignment {
	var consignments []*pb.Consignment
	for _, consignment := range c {
		consignments = append(consignments, UnMarshalConsignment(consignment))
	}
	return consignments
}

func fromStructToMap(c *core.Consignment) map[string]interface{} {
	dict := map[string]interface{}{}

	if c.Name != "" {
		dict["name"] = c.Name
	}
	if c.Description != "" {
		dict["description"] = c.Description
	}

	if len(c.Items) > 0 {
		dict["items"] = c.Items
	}

	return dict
}
