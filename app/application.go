package app

import (
	"fmt"
	"github.com/alessandroarosio/bookstore_items-api/clients/elasticsearch"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

var (
	router = mux.NewRouter()
)

func StartApplication() {
	elasticsearch.Init()

	mapUrls()
	serverAddr := "127.0.0.1:8080"

	srv := &http.Server{
		Handler:      router,
		Addr:         serverAddr,
		WriteTimeout: 500 * time.Millisecond,
		ReadTimeout:  2 * time.Millisecond,
		IdleTimeout:  60 * time.Millisecond,
	}

	fmt.Println(fmt.Sprintf("Listening on: %s", serverAddr))

	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
