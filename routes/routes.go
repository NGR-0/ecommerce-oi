package routes

import (
	"Product/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(g *gin.Engine) {
	g.POST("/users/signup", controllers.SignUp())
	g.POST("/users/login", controllers.Login())
	g.POST("/admin/addproduct", controllers.ProductViewerAdmin())
	g.GET("/users/productview", controllers.SearchProduct())
	g.GET("/users/search", controllers.SearchProductByQuery())
}
