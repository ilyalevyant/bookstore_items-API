package app

import (
	"github.com/gorilla/mux"
	"github.com/ilyalevyant/bookstore_items-API/src/clients/elasticsearch"
	"net/http"
	"time"
)

var (
	router = mux.NewRouter()
)

func StartApplication() {
	elasticsearch.Init()
	mapUrls()

	srv := &http.Server{
		Addr:         "127.0.0.1:8081",
		Handler:      router,
		ReadTimeout:  2 * time.Second,
		WriteTimeout: 1 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
