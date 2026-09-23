package http

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"

	"giparbelajar.id/belajar-rest-api-golang/internal/model"
)

// untuk dependency injection nantinya
// apabila handler perlu. eg. middleware
type BookHandler struct {
	db *gorm.DB
}

func NewBookHandler(db *gorm.DB) *BookHandler {
	return &BookHandler{
		db: db,
	}
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
	var BookInput model.Book

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

	//preparation untuk CRUD menggunakan db.[operasi CRUD]
	ctx.JSON(http.StatusOK, gin.H{
		"title":      BookInput.Title,
		"price":      BookInput.Price,
		"page_count": BookInput.PageCount,
		"rating":     BookInput.Rating,
		"descrition": BookInput.Description,
		// "Subtitle": BookInput.Subtitle,
	})
}

func (h *BookHandler) CreateBook(ctx *gin.Context) {
	var book model.Book

	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	result := h.db.Create(&book)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Buku berhasil ditambahkan",
		"data":    book,
	})
}
