package services

import (
	"github.com/Ferza17/golang_bookstore-items-api/domains/item"
	"github.com/Ferza17/golang_bookstore-items-api/logger"
	"github.com/Ferza17/golang_bookstore-items-api/utils/errors_utils"
)

var (
	ItemService itemsServiceInterface = &itemsServiceStruct{}
)

type itemsServiceInterface interface {
	Create(item.Item) (*item.Item, *errors_utils.RestError)
	Get(string) (*item.Item, *errors_utils.RestError)
}

type itemsServiceStruct struct {
}

func (s *itemsServiceStruct) Create(itemRequest item.Item) (*item.Item, *errors_utils.RestError) {
	if err := itemRequest.Save(); err != nil {
		return nil, errors_utils.NewInternalServerError("Error When trying to save to database.")
	}
	return &itemRequest, nil
}

func (s *itemsServiceStruct) Get(id string) (*item.Item, *errors_utils.RestError) {
	request := item.Item{ID: id}
	if err := request.Get(); err != nil{
		logger.Info("Error when trying get request")
		return nil, err
	}
	return &request, nil
}
