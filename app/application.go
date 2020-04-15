package app

import (
	"net/http"

	"github.com/fmarinCeiba/bookstore_items-api/clients/elasticsearch"
	"github.com/gorilla/mux"
)

var (
	router = mux.NewRouter()
)

func StartApplication() {
	elasticsearch.Init()

	mapUrls()

	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8000",
	}

	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
