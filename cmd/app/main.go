package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go-kafka-arquitetura-hexagonal /internal/infra/akafka"
	"go-kafka-arquitetura-hexagonal /internal/infra/repository"
	"go-kafka-arquitetura-hexagonal /internal/usecase"
	"go-kafka-arquitetura-hexagonal /internal/web"
	"net/http"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/go-chi/chi/v5"
	_ "github.com/go-sql-driver/mysql"
)

func routeProduct(productHandles *web.ProductHandlers) *chi.Mux {
	r := chi.NewRouter()
	r.Route("/api", func(r chi.Router) {
		r.Get("/products", productHandles.ListProducstHandle)
		r.Post("/products", productHandles.CreateProductHandle)
	})
	return r
}

func consumeProduct(createProductUseCase *usecase.CreateProductUseCase) {
	msgChan := make(chan *kafka.Message)
	go akafka.Consume([]string{"products"}, "kafka:9094", msgChan)

	for msg := range msgChan {
		dto := usecase.CreateProductInputDTO{}
		err := json.Unmarshal(msg.Value, &dto)

		if err != nil {
			fmt.Println(err)
		}

		_, err = createProductUseCase.Execute(dto)

		if err != nil {
			fmt.Println(err)
		}
	}
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(mysql:3306)/products")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	productRepository := repository.NewProductRepositoryMysql(db)
	createProductUseCase := usecase.NewCreateProductUseCase(productRepository)
	listProductsUseCase := usecase.NewListProductsUseCase(productRepository)

	productHandles := web.NewProductHandlers(createProductUseCase, listProductsUseCase)

	r := routeProduct(productHandles)

	go http.ListenAndServe(":3001", r)

	consumeProduct(createProductUseCase)
}
