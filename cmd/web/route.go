package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/ivanpahlevi8/synapsis_challange/pkg/handler"
)

func route(handler *handler.UserHandler) http.Handler {
	// create mux
	mux := chi.NewRouter()

	// create router to get id
	mux.Get("/get-user-id", handler.GetDataById)

	mux.Post("/add-user", handler.AddDataDatabase)

	mux.Put("/update-user-id", handler.UpdateDataById)

	mux.Post("/login", LoginMiddleware(http.HandlerFunc(handler.LoginHandler)))

	return mux
}
