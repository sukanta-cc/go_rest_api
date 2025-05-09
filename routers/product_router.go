package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/sukantamajhi/go_rest_api/controllers"
	"github.com/sukantamajhi/go_rest_api/middleware"
)

func ProductRouter(router *gin.RouterGroup) *gin.RouterGroup {
	productRouter := router.Group("/products")

	productRouter.POST("/", middleware.Authenticate(), controllers.CreateProduct)
	productRouter.GET("/", middleware.Authenticate(), controllers.GetProducts)
	productRouter.PUT("/:id", middleware.Authenticate(), controllers.UpdateProduct)
	productRouter.DELETE("/:id", middleware.Authenticate(), controllers.DeleteProduct)

	return productRouter
}
