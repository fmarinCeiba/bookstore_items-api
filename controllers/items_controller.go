package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/fmarinCeiba/bookstore_items-api/domain/items"
	"github.com/fmarinCeiba/bookstore_items-api/services"
	"github.com/fmarinCeiba/bookstore_items-api/utils/http_utils"
	"github.com/fmarinCeiba/bookstore_oauth-go/oauth"
	"github.com/fmarinCeiba/bookstore_utils-go/rest_errors"
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
		http_utils.ResponseError(w, *err)
		return
	}

	rBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		rErr := rest_errors.NewBadRequestError("invalid request body")
		http_utils.ResponseError(w, *rErr)
		return
	}
	defer r.Body.Close()
	var iRequest items.Item
	if err := json.Unmarshal(rBody, &iRequest); err != nil {
		rErr := rest_errors.NewBadRequestError("invalid json body")
		http_utils.ResponseError(w, *rErr)
		return
	}

	iRequest.Seller = oauth.GetCallerID(r)

	result, cErr := services.ItemsService.Create(iRequest)
	if cErr != nil {
		http_utils.ResponseError(w, *cErr)
		return
	}

	http_utils.ResponseJSON(w, http.StatusCreated, result)
}

func (c *itemsController) Get(w http.ResponseWriter, r *http.Request) {

}
