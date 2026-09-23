package main

import (
	"giparbelajar.id/belajar-rest-api-golang/internal/connection"
	"giparbelajar.id/belajar-rest-api-golang/internal/delivery/http"
	"giparbelajar.id/belajar-rest-api-golang/internal/model"
)

func main() {
	db := connection.ConnectDatabase()
	db.AutoMigrate(&model.Book{})

	bookHandler := http.NewBookHandler(db)
	router := http.SetupRouter(bookHandler)
	router.Run()
}
