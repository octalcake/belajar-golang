package http

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"giparbelajar.id/belajar-rest-api-golang/internal/model"
)

// untuk dependency injection nantinya
// apabila handler perlu. eg. middleware
type BookHandler struct{}

func NewBookHandler() *BookHandler {
	return &BookHandler{}
}

func (h *BookHandler) CheckHealth(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"messsage": "Healthy",
	})
}

func (h *BookHandler) GetMessage(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"messsage": "Hello World!",
	})
}

func (h *BookHandler) GetId(ctx *gin.Context) {
	id := ctx.Param("id") //penggunaan param untuk variable
	ctx.JSON(http.StatusOK, gin.H{"id": id})
}

func (h *BookHandler) GetTitle(ctx *gin.Context) {
	id := ctx.Param("id") //penggunaan query untuk string
	ctx.JSON(http.StatusOK, gin.H{"id": id})
}

func (h *BookHandler) PostBookHandler(ctx *gin.Context) {
	var BookInput model.BookInput

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
