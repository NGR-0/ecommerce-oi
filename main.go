package main

import (
	"Product/controllers"
	"Product/db"
	"Product/middleware"
	"Product/routes"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	app := controllers.NewApplication(db.ProductData(db.Client, "Products"), db.UserData(db.Client, "Users"))

	g := gin.New()
	g.Use(gin.Logger())

	routes.UserRoutes(g)
	g.Use(middleware.Authencation())

	g.GET("/addtocart", app.AddToCart())
	g.GET("/removeitem", app.RemoveItem())
	g.GET("/cartcheckout", app.BuyFromCart())
	g.GET("/instantbuy", app.InstantBuy())

	log.Fatal(g.Run(":" + port))
}
