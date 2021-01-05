package services

import (
	"github.com/Ferza17/golang_bookstore-items-api/domains/item"
	"github.com/Ferza17/golang_bookstore-items-api/utils/errors"
)

var (
	ItemService itemsServiceInterface = &itemsServiceStruct{}
)

type itemsServiceInterface interface {
	Create(item.Item) (*item.Item, *errors.RestError)
	Get(string) (*item.Item, *errors.RestError)
}

type itemsServiceStruct struct {
}

func (s *itemsServiceStruct) Create(item.Item) (*item.Item, *errors.RestError) {
	return nil, errors.NewInternalServerError("Implement Me")
}

func (s *itemsServiceStruct) Get(string) (*item.Item, *errors.RestError) {
	return nil, errors.NewInternalServerError("Implement Me")

}
