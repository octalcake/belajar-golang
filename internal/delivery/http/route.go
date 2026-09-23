package http

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter(BookHandler *BookHandler) *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			//anonymous function and routing
			v1.GET("/", BookHandler.CheckHealth)
			//named function and routing
			v1.GET("/message", BookHandler.GetMessage)
			//path variables
			v1.GET("/getId/:id", BookHandler.GetId)
			v1.POST("/books", BookHandler.PostBookHandler)
			v1.POST("/create-book", BookHandler.CreateBook)

		}
	}
	return router
}
