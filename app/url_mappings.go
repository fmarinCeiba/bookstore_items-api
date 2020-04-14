package app

import (
	"net/http"

	"github.com/fmarinCeiba/bookstore_items-api/controllers"
)

func mapUrls() {
	router.HandleFunc("/items", controllers.ItemsController.Create).Methods(http.MethodPost)
	router.HandleFunc("/ping", controllers.PingController.Ping).Methods(http.MethodGet)
}
