package app

import (
	"github.com/Ferza17/golang_bookstore-items-api/utils/errors"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var (
	r = mux.NewRouter()
)

func StartApplication() {
	Urls()
	srv := &http.Server{
		Handler: r,
		Addr: "127.0.0.1:8003",
	}

	if err := srv.ListenAndServe(); err != nil{
		errors.NewInternalServerError("Can't connect to the server.")
	}

	log.Println("Server running...")
}
