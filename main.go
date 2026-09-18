package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	//anonymous function and routing
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"messsage": "Healthy",
		})
	})

	//named function and routing
	router.GET("/message", getMessage)

	//path variables
	router.GET("/getId/:id", getId)
	router.POST("/books", postBookHandler)

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

	//alias JSON dalam go
	//Subtitle string `json:"sub_title"`
}

func postBookHandler(ctx *gin.Context) {
	var BookInput BookInput

	err := ctx.ShouldBindJSON(BookInput)
	if err != nil {
		//not best practice, karena server akan mati / crash
		// log.Fatal(err)
		ctx.JSON(http.StatusBadRequest, err)
		fmt.Print(err)
		return //biar ngga eksekusi code dibawahnya
	}

	ctx.JSON(http.StatusOK, gin.H{
		"title": BookInput.Title,
		"price": BookInput.Price,
		// "Subtitle": BookInput.Subtitle,
	})

}
