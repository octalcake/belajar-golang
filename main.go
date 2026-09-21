package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func main() {
	router := gin.Default()

	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			//anonymous function and routing
			v1.GET("/", func(ctx *gin.Context) {
				ctx.JSON(http.StatusOK, gin.H{
					"messsage": "Healthy",
				})
			})

			//named function and routing
			v1.GET("/message", getMessage)

			//path variables
			v1.GET("/getId/:id", getId)
			v1.POST("/books", postBookHandler)
		}
	}
	router.Run()
}

func getMessage(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"messsage": "Hello World",
	})
}

func getId(ctx *gin.Context) {
	id := ctx.Param("id") //penggunaan param untuk variable

	ctx.JSON(http.StatusOK, gin.H{"id": id})
}

func getTitle(ctx *gin.Context) {
	id := ctx.Param("id") //penggunaan query untuk string

	ctx.JSON(http.StatusOK, gin.H{"id": id})
}

type BookInput struct {
	Title string `json:"title" binding:"required,"`
	Price int    `json:"price" binding:"required,number"`

	// bisa menggunakan json.[Item] untuk menghindari internal server error
	// akibat error yang dilewati oleh go validator
	PageCount json.Number `json:"pageCount" binding:"required,number"`

	//alias JSON dalam go
	//Subtitle string `json:"sub_title"`
}

func postBookHandler(ctx *gin.Context) {
	var BookInput BookInput

	err := ctx.ShouldBindJSON(BookInput)
	if err != nil {
		//not best practice, karena server akan mati / crash
		// log.Fatal(err)
		// ctx.JSON(http.StatusBadRequest, err)
		// fmt.Print(err)
		// return //biar ngga eksekusi code dibawahnya

		for _, e := range err.(validator.ValidationErrors) {
			errMessage := fmt.Sprintf(e.Field(), e.ActualTag())
			ctx.JSON(http.StatusBadRequest, errMessage)
			return
		}

	}

	ctx.JSON(http.StatusOK, gin.H{
		"title": BookInput.Title,
		"price": BookInput.Price,
		// "Subtitle": BookInput.Subtitle,
	})

}
