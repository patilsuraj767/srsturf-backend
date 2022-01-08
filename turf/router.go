package turf

import (
	"github.com/labstack/echo/v4"
	"github.com/patilsuraj767/turf/turf/handler"
)

func SetupRoutes(e *echo.Echo) {
	api := e.Group("/api/turf")
	api.GET("/status", handler.Status)

	v1 := api.Group("/v1")

	//booking API
	booking := v1.Group("/bookings")
	booking.DELETE("/:id", handler.DeleteBooking)
	booking.GET("/:id", handler.GetBooking)
	booking.GET("/", handler.GetBookings)
	booking.POST("/", handler.CreateBooking)

	customer := v1.Group("/customers")
	customer.GET("/:id", handler.GetCustomer)
	customer.GET("/", handler.GetCustomers)
}
