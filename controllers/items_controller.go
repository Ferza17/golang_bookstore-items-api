package controllers

import (
	"encoding/json"
	"github.com/Ferza17/bookstore_oauth_library-go/oauth"
	ItemDomain "github.com/Ferza17/golang_bookstore-items-api/domains/item"
	"github.com/Ferza17/golang_bookstore-items-api/domains/queries"
	"github.com/Ferza17/golang_bookstore-items-api/services"
	restErr "github.com/Ferza17/golang_bookstore-items-api/utils/errors_utils"
	"github.com/Ferza17/golang_bookstore-items-api/utils/http_utils"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strings"
)

var (
	ItemsController itemsControllerInterface = &itemsControllerStruct{}
)

type itemsControllerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	Search(w http.ResponseWriter, r *http.Request)
}
type itemsControllerStruct struct{}

func (c *itemsControllerStruct) Create(w http.ResponseWriter, r *http.Request) {
	if err := oauth.AuthenticateRequest(r); err != nil {
		respErr := restErr.NewInternalServerError("Invalid Access Token")
		http_utils.RespondError(w, *respErr)
		return
	}

	if oauth.GetCallerId(r) == 0 {
		respErr := restErr.NewUnauthorizedError()
		http_utils.RespondError(w, *respErr)
		return
	}

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respErr := restErr.NewBadRequestError("Invalid Request Body")
		http_utils.RespondError(w, *respErr)
		return
	}
	defer r.Body.Close()

	var itemRequest ItemDomain.Item
	if err := json.Unmarshal(requestBody, &itemRequest); err != nil {
		respErr := restErr.NewBadRequestError("Invalid JSON Body")
		http_utils.RespondError(w, *respErr)
		return
	}
	itemRequest.Seller = oauth.GetCallerId(r)

	result, createErr := services.ItemService.Create(itemRequest)
	if createErr != nil {
		http_utils.RespondError(w, *createErr)
		return
	}

	http_utils.RespondJSON(w, http.StatusCreated, result)
}

func (c *itemsControllerStruct) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemId := strings.TrimSpace(vars["id"])
	item, err := services.ItemService.Get(itemId)

	if err != nil {
		respErr := restErr.NewBadRequestError("cant found data with that id")
		http_utils.RespondError(w, *respErr)
		return
	}

	http_utils.RespondJSON(w, http.StatusOK, item)

}

func (c *itemsControllerStruct) Search(w http.ResponseWriter, r *http.Request) {
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		apiError := restErr.NewBadRequestError("Invalid JSON Body")
		http_utils.RespondError(w, *apiError)
		return
	}
	defer r.Body.Close()


	var query queries.EsQuery

	if err := json.Unmarshal(bytes, &query); err != nil {
		apiError := restErr.NewBadRequestError("Error When unmarshalling body")
		http_utils.RespondError(w, *apiError)
		return
	}

	items, searchError := services.ItemService.Search(query)
	if searchError != nil {
		http_utils.RespondError(w, *searchError)
		return
	}
	http_utils.RespondJSON(w, http.StatusOK, items)
}
