package item

import (
	"encoding/json"
	"fmt"
	"github.com/Ferza17/golang_bookstore-items-api/clients/elasticsearch"
	"github.com/Ferza17/golang_bookstore-items-api/domains/queries"
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

	if len(responseBody["hits"].(map[string]interface{})["hits"].([]interface{})) == 0 {
		return restErr.NewNotFoundError("No math data with given id")
	}

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

func (i *Item) Search(query queries.EsQuery) ([]Item, *restErr.RestError) {
	finalQuery := query.Build()
	response, err := elasticsearch.Client.Search(IndexItem, finalQuery)
	if err != nil {
		return nil, restErr.NewInternalServerError("error when trying to get data")
	}
	defer response.Body.Close()
	// Get Body
	var responseBody map[string]interface{}
	if err := json.NewDecoder(response.Body).Decode(&responseBody); err != nil {
		//logger.Info("Error when decoding json")
		return nil, restErr.NewInternalServerError("Error when decoding json")
	}

	if len(responseBody["hits"].(map[string]interface{})["hits"].([]interface{})) == 0 {
		return nil, restErr.NewNotFoundError("No math data with given request")
	}

	items := make([]Item, len(responseBody["hits"].(map[string]interface{})["hits"].([]interface{})))

	for index, hit := range responseBody["hits"].(map[string]interface{})["hits"].([]interface{}) {
		doc := hit.(map[string]interface{})
		source := doc["_source"]
		var item Item
		seller, _ := strconv.ParseInt(fmt.Sprint(source.(map[string]interface{})["seller"]), 10, 64)
		item.Seller = seller

		item.Title = fmt.Sprint(source.(map[string]interface{})["title"])

		//TODO: assign value to i.Description

		item.Video = fmt.Sprint(source.(map[string]interface{})["video"])

		price, _ := strconv.ParseFloat(fmt.Sprint(source.(map[string]interface{})["price"]), 32)
		item.Price = float32(price)

		AvailableQuantity, _ := strconv.ParseInt(fmt.Sprint(source.(map[string]interface{})["available_quantity"]), 10, 64)
		item.AvailableQuantity = AvailableQuantity

		SoldQuantity, _ := strconv.ParseInt(fmt.Sprint(source.(map[string]interface{})["sold_quantity"]), 10, 64)
		item.SoldQuantity = SoldQuantity

		item.Status = fmt.Sprint(source.(map[string]interface{})["status"])

		items[index] = item
	}

	return items, nil
}
