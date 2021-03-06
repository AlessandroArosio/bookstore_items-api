package app

import (
	"github.com/alessandroarosio/bookstore_items-api/controllers"
	"net/http"
)

func mapUrls() {
	router.HandleFunc("/items", controllers.ItemsController.Create).Methods(http.MethodPost)
	router.HandleFunc("/", controllers.PingController.Pong).Methods(http.MethodGet)
}
