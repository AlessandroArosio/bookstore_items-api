package controllers

import (
	"encoding/json"
	"github.com/alessandroarosio/bookstore-utils-go/rest_errors"
	"github.com/alessandroarosio/bookstore_items-api/domain/items"
	"github.com/alessandroarosio/bookstore_items-api/services"
	"github.com/alessandroarosio/bookstore_items-api/utils/http_utils"
	"github.com/alessandroarosio/bookstore_oauth-go/oauth"
	"io/ioutil"

	"net/http"
)

var (
	ItemsController itemsControllerInterface = &itemsController{}
)

type itemsControllerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
}

type itemsController struct{}

func (c *itemsController) Create(w http.ResponseWriter, r *http.Request) {
	if err := oauth.AuthenticateRequest(r); err != nil {
		http_utils.RespondError(w, err)
		return
	}
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respErr := rest_errors.NewBadRequestError("invalid request body")
		http_utils.RespondError(w, respErr)
		return
	}
	defer r.Body.Close()

	var itemRequest items.Item
	if err := json.Unmarshal(requestBody, &itemRequest); err != nil {
		respErr := rest_errors.NewBadRequestError("invalid json body")
		http_utils.RespondError(w, respErr)
		return
	}

	itemRequest.Seller = oauth.GetCallerId(r)

	result, err := services.ItemsService.Create(itemRequest)
	if err != nil {
		mapErr := rest_errors.NewBadRequestError("error when creating item")
		http_utils.RespondError(w, mapErr)
		return
	}

	http_utils.RespondJson(w, http.StatusCreated, result)
}

func (c *itemsController) Get(w http.ResponseWriter, r *http.Request) {

}
