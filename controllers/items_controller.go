package controllers

import (
	"fmt"
	"github.com/Ferza17/bookstore_oauth_library-go/oauth"
	ItemDomain "github.com/Ferza17/golang_bookstore-items-api/domains/item"
	"github.com/Ferza17/golang_bookstore-items-api/services"
	"net/http"
)

var (
	ItemsController itemsControllerInterface = &itemsControllerStruct{}
)

type itemsControllerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
}
type itemsControllerStruct struct{}

func (c *itemsControllerStruct) Create(w http.ResponseWriter, r *http.Request) {
	if err := oauth.AuthenticateRequest(r); err != nil {
		// TODO: Return error to the caller
		return
	}
	item := ItemDomain.Item{
		Seller: oauth.GetCallerId(r),
	}
	result, err := services.ItemService.Create(item)
	if err != nil {
		//TODO: Return error json to the user
		return
	}

	fmt.Println(result)
	//TODO: Return created item as JSON with  http status 201: Created
}

func (c *itemsControllerStruct) Get(w http.ResponseWriter, r *http.Request) {

}
