package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
)

func main() {
	log.Println("Starting the order microservice")

	ctx := cmd.Context()

	r, closeFn := createOrderMicroservice()

	defer closeFn()

	server := &http.Server{Addr: os.Getenv("SHOP_ORDER_SERVICE_BIND_ADDR"), Handler: r}

	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			panic(err)
		}
	}()

	<-ctx.Done()

	log.Println("closing order microservice")

	if err := server.Close(); err != nil {
		panic(err)
	}
}

func createOrderMicroservice() (router *chi.Mux, closeFn func()) {
	cmd.WaitForService(os.Getenv("SHOP_RABBITMQ_ADDR"))

	shopHTTPClient := orders_infra_product.NewHTTPClient(os.Getenv("SHOP_SERVICE_ADDR"))

	orders_public_http.AddRoutes(r, oderseService, ordersRepo)
	orders_private_http.AddRoutes(r, ordersService, ordersRepo)

	return r, func() {

	}
}
