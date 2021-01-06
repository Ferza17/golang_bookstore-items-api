package app

import (
	"github.com/Ferza17/golang_bookstore-items-api/clients/elasticsearch"
	"github.com/Ferza17/golang_bookstore-items-api/utils/errors_utils"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

var (
	r = mux.NewRouter()
)

func StartApplication() {
	elasticsearch.Init()
	Urls()
	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8083",
		WriteTimeout: 500 * time.Millisecond,
		ReadTimeout:  2 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
	r.Headers("Content-Type", "application/json")
	if err := srv.ListenAndServe(); err != nil {
		errors_utils.NewInternalServerError("Can't connect to the server.")
	}
}
