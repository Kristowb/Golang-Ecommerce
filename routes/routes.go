package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kristowb/ecommerce-backend/controllers"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("users/signup", controllers.signUp())
	incomingRoutes.POST("users/login", controllers.login())
	incomingRoutes.POST("/admin/addproduct", controllers.ProductViewerAdmin())
	incomingRoutes.GET("/users/productview", controllers.SearchProduct())
	incomingRoutes.GET("/users/search", controllers.SearchProductByQuery())
}
