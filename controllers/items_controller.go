package controllers

import (
	"fmt"
	"net/http"

	"github.com/fmarinCeiba/bookstore_items-api/domain/items"
	"github.com/fmarinCeiba/bookstore_items-api/services"
	"github.com/fmarinCeiba/bookstore_oauth-go/oauth"
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
		return
	}

	item := items.Item{
		Seller: oauth.GetCallerID(r),
	}

	result, err := services.ItemsService.Create(item)
	if err != nil {

	}

	fmt.Println(result)

}

func (c *itemsController) Get(w http.ResponseWriter, r *http.Request) {

}
