package item

import (
	"encoding/json"
	"fmt"
	"github.com/Ferza17/golang_bookstore-items-api/clients/elasticsearch"
	"github.com/Ferza17/golang_bookstore-items-api/logger"
	restErr "github.com/Ferza17/golang_bookstore-items-api/utils/errors_utils"
	"strconv"
	"strings"
)

const (
	IndexItem = "items"
)

func (i *Item) Save() *restErr.RestError {
	result, err := elasticsearch.Client.Index(IndexItem, i)
	if err != nil {
		logger.Error(fmt.Sprintf("error when trying to index document in index %s", IndexItem), err)
		return restErr.NewInternalServerError("error when trying to index document to es")
	}
	defer result.Body.Close()
	//GET ID
	var responseBody map[string]interface{}
	if err := json.NewDecoder(result.Body).Decode(&responseBody); err != nil {
		return restErr.NewInternalServerError("Error when decoding json")
	}
	i.ID = fmt.Sprintf("%v", responseBody["_id"])
	return nil
}

func (i *Item) Get() *restErr.RestError {
	result, err := elasticsearch.Client.Get(IndexItem, i.ID)
	if err != nil {
		if strings.Contains(err.Error(), "404") {
			return restErr.NewNotFoundError("no item found with that id")
		}
		logger.Error(fmt.Sprintf("error when trying to get id : %s", i.ID), err)
		return restErr.NewInternalServerError(fmt.Sprintf("error when trying to get id : %s", i.ID))
	}
	defer result.Body.Close()
	// Get Body
	var responseBody map[string]interface{}
	if err := json.NewDecoder(result.Body).Decode(&responseBody); err != nil {
		//logger.Info("Error when decoding json")
		return restErr.NewInternalServerError("Error when decoding json")
	}
	logger.Info(fmt.Sprintln(responseBody["hits"]))
	for _, hit := range responseBody["hits"].(map[string]interface{})["hits"].([]interface{}) {
		doc := hit.(map[string]interface{})
		source := doc["_source"]
		seller, _ := strconv.ParseInt(fmt.Sprint(source.(map[string]interface{})["seller"]), 10, 64)

		i.Seller = seller
		i.Title = fmt.Sprint(source.(map[string]interface{})["title"])

		//TODO: assign value to i.Description

		i.Video = fmt.Sprint(source.(map[string]interface{})["video"])

		price, _ := strconv.ParseFloat(fmt.Sprint(source.(map[string]interface{})["price"]), 32)
		i.Price = float32(price)

		AvailableQuantity, _ := strconv.ParseInt(fmt.Sprint(source.(map[string]interface{})["available_quantity"]), 10, 64)
		i.AvailableQuantity = AvailableQuantity

		SoldQuantity, _ := strconv.ParseInt(fmt.Sprint(source.(map[string]interface{})["sold_quantity"]), 10, 64)
		i.SoldQuantity = SoldQuantity

		i.Status = fmt.Sprint(source.(map[string]interface{})["status"])
	}
	return nil
}
