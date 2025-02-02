package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/LambdaIITH/Dashboard/backend/internal/controller"
	"github.com/LambdaIITH/Dashboard/backend/internal/middlewares"
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
		authGroup.GET("/logout", controller.LogoutHandler)
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

	userGroup := router.Group("/user")
	{
		userGroup.GET("/", controller.User)
		userGroup.PATCH("/update", controller.UpdateUser)
		userGroup.PATCH("/update/fcm", controller.UpdateUserFCMToken)
	}

	router.POST("found/add_item", controller.AddFoundItemHandler)
	router.GET("/found/all", controller.GetAllFoundItemsHandler)
	router.GET("/found/get_item/:id", controller.GetFoundItemByIdHandler)
	router.PUT("/found/edit_item", controller.EditFoundItemHandler)
	router.POST("/found/delete_item", controller.DeleteFoundItemHandler)
	router.GET("/found/search", controller.SearchFoundItemHandler)

	//Group routes for timetable/calendar
	timetableGroup := router.Group("/schedule")
	{
		timetableGroup.GET("/all_courses", middlewares.AuthMiddleware(), controller.GetAllCourses)
		timetableGroup.GET("/courses", middlewares.AuthMiddleware(), controller.GetTimetable)
		timetableGroup.POST("/courses", middlewares.AuthMiddleware(), controller.PostEditTimetable)
		timetableGroup.GET("/share/:code", middlewares.AuthMiddleware(), controller.GetSharedTimetable)
		timetableGroup.POST("/share", middlewares.AuthMiddleware(), controller.PostSharedTimetable)
		timetableGroup.DELETE("/share/:code", middlewares.AuthMiddleware(), controller.DeleteSharedTimetable)
	}

	// GET : /announcements?limit=4&offset=4
	router.GET("/announcements", controller.GetAnnouncements)
	router.Static("/announcements/images", "announcementImages/")
	router.POST("/announcements", controller.PostAnnouncement)

	cabshareGroup := router.Group("/cabshare")
	{
		cabshareGroup.GET("/me", controller.CheckAuth)
		cabshareGroup.POST("/bookings", controller.CreateBooking)
		cabshareGroup.PATCH("/bookings/:booking_id", controller.UpdateBooking)
		cabshareGroup.GET("/me/bookings", controller.UserBookings)
		cabshareGroup.GET("/me/requests", controller.UserRequests)
		cabshareGroup.GET("/bookings", controller.SearchBookings)
		cabshareGroup.POST("/bookings/:booking_id/request", controller.RequestToJoinBooking)
		cabshareGroup.DELETE("/bookings/:booking_id/request", controller.DeleteRequest)
		cabshareGroup.POST("/bookings/:booking_id/accept", controller.AcceptRequest)
		cabshareGroup.POST("/bookings/:booking_id/reject", controller.RejectRequest)
		cabshareGroup.DELETE("/bookings/:booking_id", controller.DeleteExistingBooking)
		cabshareGroup.DELETE("/bookings/:booking_id/self", controller.ExitBooking)
	}
}
