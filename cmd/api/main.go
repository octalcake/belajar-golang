package main

import (
	"giparbelajar.id/belajar-rest-api-golang/internal/delivery/http"
)

func main() {
	bookHandler := http.NewBookHandler()
	router := http.SetupRouter(bookHandler)
	router.Run()
}
