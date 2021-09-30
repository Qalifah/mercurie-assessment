// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/shipping.proto

package shipping

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/asim/go-micro/v3/api"
	client "github.com/asim/go-micro/v3/client"
	server "github.com/asim/go-micro/v3/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for ShippingService service

func NewShippingServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for ShippingService service

type ShippingService interface {
	CreateConsignment(ctx context.Context, in *Consignment, opts ...client.CallOption) (*Response, error)
	GetConsignment(ctx context.Context, in *SearchParameter, opts ...client.CallOption) (*Response, error)
	GetAllConsignments(ctx context.Context, in *GetAllRequest, opts ...client.CallOption) (*Response, error)
	UpdateConsignment(ctx context.Context, in *Consignment, opts ...client.CallOption) (*Response, error)
	DeleteConsignment(ctx context.Context, in *SearchParameter, opts ...client.CallOption) (*Response, error)
	QuoteConsignment(ctx context.Context, in *SearchParameter, opts ...client.CallOption) (*Response, error)
}

type shippingService struct {
	c    client.Client
	name string
}

func NewShippingService(name string, c client.Client) ShippingService {
	return &shippingService{
		c:    c,
		name: name,
	}
}

func (c *shippingService) CreateConsignment(ctx context.Context, in *Consignment, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "ShippingService.CreateConsignment", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shippingService) GetConsignment(ctx context.Context, in *SearchParameter, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "ShippingService.GetConsignment", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shippingService) GetAllConsignments(ctx context.Context, in *GetAllRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "ShippingService.GetAllConsignments", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shippingService) UpdateConsignment(ctx context.Context, in *Consignment, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "ShippingService.UpdateConsignment", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shippingService) DeleteConsignment(ctx context.Context, in *SearchParameter, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "ShippingService.DeleteConsignment", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shippingService) QuoteConsignment(ctx context.Context, in *SearchParameter, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "ShippingService.QuoteConsignment", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ShippingService service

type ShippingServiceHandler interface {
	CreateConsignment(context.Context, *Consignment, *Response) error
	GetConsignment(context.Context, *SearchParameter, *Response) error
	GetAllConsignments(context.Context, *GetAllRequest, *Response) error
	UpdateConsignment(context.Context, *Consignment, *Response) error
	DeleteConsignment(context.Context, *SearchParameter, *Response) error
	QuoteConsignment(context.Context, *SearchParameter, *Response) error
}

func RegisterShippingServiceHandler(s server.Server, hdlr ShippingServiceHandler, opts ...server.HandlerOption) error {
	type shippingService interface {
		CreateConsignment(ctx context.Context, in *Consignment, out *Response) error
		GetConsignment(ctx context.Context, in *SearchParameter, out *Response) error
		GetAllConsignments(ctx context.Context, in *GetAllRequest, out *Response) error
		UpdateConsignment(ctx context.Context, in *Consignment, out *Response) error
		DeleteConsignment(ctx context.Context, in *SearchParameter, out *Response) error
		QuoteConsignment(ctx context.Context, in *SearchParameter, out *Response) error
	}
	type ShippingService struct {
		shippingService
	}
	h := &shippingServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&ShippingService{h}, opts...))
}

type shippingServiceHandler struct {
	ShippingServiceHandler
}

func (h *shippingServiceHandler) CreateConsignment(ctx context.Context, in *Consignment, out *Response) error {
	return h.ShippingServiceHandler.CreateConsignment(ctx, in, out)
}

func (h *shippingServiceHandler) GetConsignment(ctx context.Context, in *SearchParameter, out *Response) error {
	return h.ShippingServiceHandler.GetConsignment(ctx, in, out)
}

func (h *shippingServiceHandler) GetAllConsignments(ctx context.Context, in *GetAllRequest, out *Response) error {
	return h.ShippingServiceHandler.GetAllConsignments(ctx, in, out)
}

func (h *shippingServiceHandler) UpdateConsignment(ctx context.Context, in *Consignment, out *Response) error {
	return h.ShippingServiceHandler.UpdateConsignment(ctx, in, out)
}

func (h *shippingServiceHandler) DeleteConsignment(ctx context.Context, in *SearchParameter, out *Response) error {
	return h.ShippingServiceHandler.DeleteConsignment(ctx, in, out)
}

func (h *shippingServiceHandler) QuoteConsignment(ctx context.Context, in *SearchParameter, out *Response) error {
	return h.ShippingServiceHandler.QuoteConsignment(ctx, in, out)
}
