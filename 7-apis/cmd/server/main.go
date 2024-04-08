package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/tallyto/curso-go/7-apis/configs"
	"github.com/tallyto/curso-go/7-apis/internal/entity"
	"github.com/tallyto/curso-go/7-apis/internal/infra/database"
	"github.com/tallyto/curso-go/7-apis/internal/infra/webserver/handlers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	_, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.Product{}, &entity.User{})
	productDB := database.NewProduct(db)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	productHandler := handlers.NewProductHandler(productDB)

	r.Post("/products", productHandler.CreateProduct)

	// http.HandleFunc("/products", productHandler.CreateProduct)

	http.ListenAndServe(":3001", r)

}
