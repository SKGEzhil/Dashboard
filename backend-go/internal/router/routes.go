package router

import (
	"net/http"

	"github.com/LambdaIITH/Dashboard/backend/internal/controller"
	"github.com/gin-gonic/gin"
)

func home(c *gin.Context) {
	HTMLString := "<h1>Hello from <a href='https://iith.dev' target='_blank'>Lambda IITH</a></h1>"
	c.Writer.WriteHeader(http.StatusOK)

	c.Writer.Write([]byte(HTMLString))
}

func SetupRoutes(router *gin.Engine) {
	
	// Home route
	router.GET("/", home)

	// Group routes for authentication
	authGroup := router.Group("/auth")
	{
		authGroup.POST("/login", controller.LoginHandler)
		authGroup.POST("/logout", controller.LogoutHandler)
	}

	// Group routes for lost items
	lostGroup := router.Group("/lost")
	{
		lostGroup.POST("/add_item", controller.AddItemHandler)
		lostGroup.GET("/all", controller.GetAllItemsHandler)
		lostGroup.GET("/get_item/:id", controller.GetItemByIdHandler)
		lostGroup.PUT("/edit_item", controller.EditItemHandler)
		lostGroup.POST("/delete_item", controller.DeleteItemHandler)
		lostGroup.GET("/search", controller.SearchItemHandler)
	}
	
	// Group routes for transport
	transportGroup := router.Group("/transport")
	{
		transportGroup.GET("/", controller.GetBusSchedule)
		transportGroup.GET("/cityBus", controller.GetCityBusSchedule)
		transportGroup.POST("/qr", controller.ProcessTransaction)
		transportGroup.POST("/qr/scan", controller.ScanQRCode)
	}

	sellGroup := router.Group("/sell")
	{
		sellGroup.POST("/add_item", controller.AddSellItemHandler)
		sellGroup.GET("/all", controller.GetAllSellItemsHandler)
		sellGroup.GET("/get_item/:id", controller.GetSellItemByIdHandler)
		sellGroup.PUT("/edit_item", controller.EditSellItemHandler)
		sellGroup.POST("/delete_item", controller.DeleteSellItemHandler)
		sellGroup.GET("/search", controller.SearchSellItemHandler)
	}

}
