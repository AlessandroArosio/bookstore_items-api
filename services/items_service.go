package services

import (
	"github.com/alessandroarosio/bookstore-utils-go/rest_errors"
	"github.com/alessandroarosio/bookstore_items-api/domain/items"
	"net/http"
)

var (
	ItemsService itemsServiceInterface = &itemService{}
)

type itemsServiceInterface interface {
	Create(item items.Item) (*items.Item, rest_errors.RestErr)
	Get(id string) (*items.Item, rest_errors.RestErr)
}

type itemService struct{}

func (s *itemService) Create(item items.Item) (*items.Item, rest_errors.RestErr) {
	return nil, rest_errors.NewRestError("todo", http.StatusNotImplemented, "not_implemented", nil)
}

func (s *itemService) Get(id string) (*items.Item, rest_errors.RestErr) {
	return nil, rest_errors.NewRestError("todo", http.StatusNotImplemented, "not_implemented", nil)
}
