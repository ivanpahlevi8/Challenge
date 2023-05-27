package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/ivanpahlevi8/synapsis_challange/pkg/handler"
)

func route(handler *handler.UserHandler, itemHandler *handler.ItemHandler, userActivityHandler *handler.UserActivityHandler) http.Handler {
	// create mux
	mux := chi.NewRouter()

	// middleware
	mux.Use(SessionMiddleware)

	// index
	mux.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		w.Write([]byte("Welcome To Home Page"))

		appConfig.Session.Put(r.Context(), "username", "not logged")
	})

	// create router to get id
	mux.Get("/get-user-id", handler.GetDataById)

	mux.Post("/add-user", handler.AddDataDatabase)

	mux.Put("/update-user-id", handler.UpdateDataById)

	mux.Post("/login", LoginMiddleware(http.HandlerFunc(handler.LoginHandler)))

	// create route for item
	mux.Get("/get-item-id", itemHandler.GetDataById)

	mux.Post("/add-item", itemHandler.AddItem)

	// create route for user activiy
	mux.Put("/user-get-item", ActiveUserMiddleware(http.HandlerFunc(userActivityHandler.UserAddItem)))

	mux.Put("/user-delete-item", ActiveUserMiddleware(http.HandlerFunc(userActivityHandler.UserRemoveItem)))

	mux.Get("/user-get-all-item", ActiveUserMiddleware(http.HandlerFunc(userActivityHandler.UserViewAllItemInChart)))

	mux.Get("/user-checkout", ActiveUserMiddleware(http.HandlerFunc(userActivityHandler.CheckoutAllItem)))

	mux.Get("/user-category", ActiveUserMiddleware(http.HandlerFunc(userActivityHandler.GetAllItemsBasedOnCategory)))

	return mux
}
