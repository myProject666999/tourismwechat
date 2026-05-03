package router

import (
	"tourismwechat/controllers"
	"tourismwechat/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	api := r.Group("/api")
	{
		api.POST("/register", controllers.Register)
		api.POST("/login", controllers.Login)
		api.POST("/admin/login", controllers.AdminLogin)

		api.GET("/announcements", controllers.GetAnnouncementList)
		api.GET("/announcements/:id", controllers.GetAnnouncementDetail)
		api.GET("/scenics", controllers.GetScenicList)
		api.GET("/scenics/:id", controllers.GetScenicDetail)
		api.GET("/hotels", controllers.GetHotelList)
		api.GET("/hotels/:id", controllers.GetHotelDetail)
		api.GET("/routes", controllers.GetRouteList)
		api.GET("/routes/:id", controllers.GetRouteDetail)
		api.GET("/daily-recommendations", controllers.GetDailyRecommendations)

		api.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"code":    200,
				"message": "pong",
			})
		})
	}

	user := api.Group("/user").Use(middleware.JWTAuth())
	{
		user.GET("/profile", controllers.GetProfile)
		user.PUT("/profile", controllers.UpdateProfile)

		user.POST("/announcements/:id/like", controllers.LikeAnnouncement)
		user.POST("/announcements/:id/dislike", controllers.DislikeAnnouncement)
		user.POST("/scenics/:id/like", controllers.LikeScenic)
		user.POST("/scenics/:id/dislike", controllers.DislikeScenic)

		user.POST("/bookings", controllers.CreateBooking)
		user.GET("/bookings", controllers.GetUserBookings)
		user.GET("/bookings/:id", controllers.GetBookingDetail)
		user.POST("/bookings/:id/cancel", controllers.CancelBooking)

		user.POST("/itineraries", controllers.CreateItinerary)
		user.GET("/itineraries", controllers.GetUserItineraries)
		user.GET("/itineraries/:id", controllers.GetItineraryDetail)
		user.POST("/itinerary-items", controllers.AddItineraryItem)
		user.PUT("/itinerary-items/:id", controllers.UpdateItineraryItem)
		user.DELETE("/itinerary-items/:id", controllers.DeleteItineraryItem)
	}

	admin := api.Group("/admin").Use(middleware.JWTAuth(), middleware.AdminAuth())
	{
		admin.GET("/dashboard", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"code":    200,
				"message": "success",
				"data": gin.H{
					"total_users":    0,
					"total_hotels":   0,
					"total_bookings": 0,
					"total_revenue":  0,
				},
			})
		})

		admin.GET("/admins", controllers.GetAdminList)
		admin.POST("/admins", controllers.CreateAdmin)
		admin.PUT("/admins/:id", controllers.UpdateAdmin)
		admin.DELETE("/admins/:id", controllers.DeleteAdmin)

		admin.GET("/users", controllers.GetUserList)
		admin.PUT("/users/:id/status", controllers.UpdateUserStatus)
		admin.PUT("/users/:id", controllers.UpdateUser)
		admin.DELETE("/users/:id", controllers.DeleteUser)
		admin.POST("/users", controllers.CreateUser)

		admin.GET("/announcements", controllers.AdminGetAnnouncementList)
		admin.POST("/announcements", controllers.CreateAnnouncement)
		admin.PUT("/announcements/:id", controllers.UpdateAnnouncement)
		admin.DELETE("/announcements/:id", controllers.DeleteAnnouncement)

		admin.GET("/scenics", controllers.AdminGetScenicList)
		admin.POST("/scenics", controllers.CreateScenic)
		admin.PUT("/scenics/:id", controllers.UpdateScenic)
		admin.DELETE("/scenics/:id", controllers.DeleteScenic)

		admin.GET("/hotels", controllers.AdminGetHotelList)
		admin.POST("/hotels", controllers.CreateHotel)
		admin.PUT("/hotels/:id", controllers.UpdateHotel)
		admin.DELETE("/hotels/:id", controllers.DeleteHotel)

		admin.GET("/routes", controllers.AdminGetRouteList)
		admin.POST("/routes", controllers.CreateRoute)
		admin.PUT("/routes/:id", controllers.UpdateRoute)
		admin.DELETE("/routes/:id", controllers.DeleteRoute)

		admin.GET("/bookings", controllers.AdminGetBookingList)
		admin.PUT("/bookings/:id/status", controllers.UpdateBookingStatus)

		admin.GET("/itineraries", controllers.AdminGetItineraryList)
		admin.PUT("/itineraries/:id/status", controllers.AdminUpdateItineraryStatus)

		admin.GET("/daily-recommendations", controllers.AdminGetDailyRecommendList)
		admin.POST("/daily-recommendations", controllers.CreateDailyRecommend)
		admin.PUT("/daily-recommendations/:id", controllers.UpdateDailyRecommend)
		admin.DELETE("/daily-recommendations/:id", controllers.DeleteDailyRecommend)

		admin.GET("/hotel-stats", controllers.GetHotelStats)
		admin.GET("/hotel-stats/summary", controllers.GetHotelStatsSummary)
		admin.POST("/hotel-stats/generate", controllers.GenerateHotelStat)

		admin.GET("/export/users", controllers.ExportUsers)
		admin.POST("/import/users", controllers.ImportUsers)
		admin.GET("/export/hotels", controllers.ExportHotels)
		admin.GET("/export/bookings", controllers.ExportBookings)
	}

	return r
}
