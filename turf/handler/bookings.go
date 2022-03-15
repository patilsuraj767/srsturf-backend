package handler

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/patilsuraj767/turf/turf/model"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

func EditBooking(c echo.Context) error {
	var booking model.Booking
	err := booking.GetBookingsById(c.Param("id"))
	if err != nil {
		return c.JSON(500, echo.Map{"status": "error", "message": err})
	}

	payload := struct {
		Name        string `json:"name"`
		Mobile      string `json:"mobile"`
		BookingDate string `json:"bookingDate"`
		StartTime   string `json:"startTime"`
		EndTime     string `json:"endTime"`
		Description string `json:"description"`
	}{}
	if err := c.Bind(&payload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var customer model.Customer
	// Create customer in database if it does not exist.
	if err := customer.GetCustomerByMobile(payload.Mobile); err.Error != nil {
		if errors.Is(err.Error, gorm.ErrRecordNotFound) {
			customer.Name = payload.Name
			customer.Mobile = payload.Mobile
			result := customer.CreateCustomer()
			if result.Error != nil {
				return c.JSON(500, echo.Map{"status": "error", "message": "Error Created Booking Account"})
			}
		} else {
			fmt.Println("Some database issue", err.Error)
		}
	}

	// Make booking in database.
	t := time.Now()
	currentdayBegins := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	bookingDate, _ := time.Parse("2006-1-02", payload.BookingDate)
	startTime, _ := time.Parse("2006-1-02 15:04:05", payload.BookingDate+" "+payload.StartTime)
	endTime, _ := time.Parse("2006-1-02 15:04:05", payload.BookingDate+" "+payload.EndTime)

	if currentdayBegins.After(bookingDate) {
		return c.JSON(500, echo.Map{"status": "error", "message": "Cannot do past bookings"})
	}
	if startTime.After(endTime) {
		return c.JSON(500, echo.Map{"status": "error", "message": "Invalid time. startTime can not be > endTime"})
	}

	booking.Customer = customer
	booking.BookingDate = datatypes.Date(bookingDate)
	booking.StartTime = datatypes.NewTime(startTime.Hour(), startTime.Minute(), startTime.Second(), 0)
	booking.EndTime = datatypes.NewTime(endTime.Hour(), endTime.Minute(), endTime.Second(), 0)
	booking.Description = payload.Description

	result, err := booking.EditBooking()
	if err != nil {
		return c.JSON(500, echo.Map{"status": "error", "message": err.Error()})
	} else if result.Error != nil {
		return c.JSON(500, echo.Map{"status": "error", "message": "Error Created Booking"})
	} else {
		return c.JSON(201, echo.Map{"status": "success", "message": "Created Booking", "data": payload})
	}
}

func DeleteBooking(c echo.Context) error {
	var booking model.Booking
	tmp, _ := strconv.Atoi(c.Param("id"))
	booking.ID = uint(tmp)
	err := booking.DeleteBookingById()
	if err != nil {
		return c.JSON(500, echo.Map{"status": "error", "message": err})
	} else {
		return c.JSON(200, booking)
	}
}

func GetBooking(c echo.Context) error {
	var booking model.Booking
	err := booking.GetBookingsById(c.Param("id"))
	if err != nil {
		return c.JSON(500, echo.Map{"status": "error", "message": err})
	} else {
		return c.JSON(200, booking)
	}
}

func GetBookings(c echo.Context) error {
	t := time.Now()
	var booking model.Booking
	bookings, err := booking.GetBookingsByDate(t)
	if err != nil {
		return c.JSON(500, echo.Map{"status": "error", "message": err})
	} else {
		return c.JSON(200, bookings)
	}
}

func GetBookingsByDate(c echo.Context) error {
	location, _ := time.LoadLocation("Asia/Calcutta")
	bookingDate, _ := time.ParseInLocation("2006-1-02", c.Param("date"), location)
	var booking model.Booking
	bookings, err := booking.GetBookingsByDate(bookingDate)
	if err != nil {
		return c.JSON(500, echo.Map{"status": "error", "message": err})
	} else {
		return c.JSON(200, bookings)
	}
}

func CreateBooking(c echo.Context) error {
	var customer model.Customer
	payload := struct {
		Name        string `json:"name"`
		Mobile      string `json:"mobile"`
		BookingDate string `json:"bookingDate"`
		StartTime   string `json:"startTime"`
		EndTime     string `json:"endTime"`
		Description string `json:"description"`
	}{}
	if err := c.Bind(&payload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Create customer in database if it does not exist.
	if err := customer.GetCustomerByMobile(payload.Mobile); err.Error != nil {
		if errors.Is(err.Error, gorm.ErrRecordNotFound) {
			customer.Name = payload.Name
			customer.Mobile = payload.Mobile
			result := customer.CreateCustomer()
			if result.Error != nil {
				return c.JSON(500, echo.Map{"status": "error", "message": "Error Created Booking Account"})
			}
		} else {
			fmt.Println("Some database issue", err.Error)
		}
	}

	// Make booking in database.
	t := time.Now()
	currentdayBegins := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	bookingDate, _ := time.Parse("2006-1-02", payload.BookingDate)
	startTime, _ := time.Parse("2006-1-02 15:04", payload.BookingDate+" "+payload.StartTime)
	endTime, _ := time.Parse("2006-1-02 15:04", payload.BookingDate+" "+payload.EndTime)

	if currentdayBegins.After(bookingDate) {
		return c.JSON(500, echo.Map{"status": "error", "message": "Cannot do past bookings"})
	}
	if startTime.After(endTime) {
		return c.JSON(500, echo.Map{"status": "error", "message": "Invalid time. startTime can not be > endTime"})
	}

	var booking = model.Booking{
		Customer:    customer,
		BookingDate: datatypes.Date(bookingDate),
		StartTime:   datatypes.NewTime(startTime.Hour(), startTime.Minute(), startTime.Second(), 0),
		EndTime:     datatypes.NewTime(endTime.Hour(), endTime.Minute(), endTime.Second(), 0),
		Description: payload.Description,
	}

	result, err := booking.CreateBooking()

	if err != nil {
		return c.JSON(500, echo.Map{"status": "error", "message": err.Error()})
	} else if result.Error != nil {
		return c.JSON(500, echo.Map{"status": "error", "message": "Error Created Booking"})
	} else {
		return c.JSON(201, echo.Map{"status": "success", "message": "Created Booking", "data": payload})
	}
}
