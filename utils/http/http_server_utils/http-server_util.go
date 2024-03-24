package http_server_utils

import (
	product_handlers "api/handlers/products_handlers"
	"api/utils/logger_utils"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// var routes = []

func LaunchHttpServer() {
	port := os.Getenv("PORT")

	server := http.NewServeMux()

	logger := logger_utils.NewLogger()

	// Handlers
	productHandler := product_handlers.NewProductHandler(logger)

	// Assign Routes to handlers
	server.Handle("/product/", productHandler)

	// http server configs
	httpServer := http.Server{
		Addr:         ":" + port,
		Handler:      server,
		IdleTimeout:  time.Second,
		ReadTimeout:  time.Second,
		WriteTimeout: time.Second,
	}

	fmt.Println("Server is on Fire on Port", port)

	if err := httpServer.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
