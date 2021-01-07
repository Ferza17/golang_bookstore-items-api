package elasticsearch

import (
	"context"
	"github.com/Ferza17/golang_bookstore-items-api/logger"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/elastic/go-elasticsearch/v7/esutil"
	"io"
	"net/http"
	"time"
)

var (
	Client esClientInterface = &esClientStruct{}
)

type esClientStruct struct {
	client *elasticsearch.Client
}

type esClientInterface interface {
	setClient(*elasticsearch.Client)
	Index(string, interface{}) (*esapi.Response, error)
	Get(string, string) (*esapi.Response, error)
	Search(string, io.Reader) (*esapi.Response, error)
}

func Init() {
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
		},
		Transport: &http.Transport{
			ResponseHeaderTimeout: 10 * time.Second,
		},
	}
	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	logger.Info("Connection Successfully")
	Client.setClient(client)
}

func (c *esClientStruct) setClient(client *elasticsearch.Client) {
	c.client = client
}
func (c *esClientStruct) Index(index string, item interface{}) (*esapi.Response, error) {
	res, err := c.client.Index(index, esutil.NewJSONReader(item))
	if err != nil {
		return nil, err
	}
	return res, nil

}

func (c *esClientStruct) Get(index string, id string) (*esapi.Response, error) {
	// Preparing Request
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"_id": id,
			},
		},
	}

	res, err := c.client.Search(
		c.client.Search.WithContext(context.Background()),
		c.client.Search.WithIndex(index),
		c.client.Search.WithPretty(),
		c.client.Search.WithBody(esutil.NewJSONReader(query)),
	)

	if err != nil {
		return nil, err
	}

	if res.IsError() {
		return nil, err
	}
	return res, nil
}

func (c *esClientStruct) Search(index string, body io.Reader) (*esapi.Response, error) {
	//TODO: fix error here
	res, err := c.client.Search(
		c.client.Search.WithContext(context.Background()),
		c.client.Search.WithIndex(index),
		c.client.Search.WithPretty(),
		c.client.Search.WithBody(body),
	)
	if err != nil {
		return nil, err
	}
	if res.IsError() {
		return nil, err
	}

	return res, nil
}
