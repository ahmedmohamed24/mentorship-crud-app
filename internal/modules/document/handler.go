package document

import "github.com/gin-gonic/gin"

func RegisterHandlers(router *gin.RouterGroup, service Service) {
	router.POST("/", service.CreateDocument)
	router.PUT("/:id", service.UpdateDocument)
	router.DELETE("/:id", service.DeleteDocument)
	router.GET("/:id", service.GetDocument)
	router.GET("/", service.GetAllDocuments)
}
