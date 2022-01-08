package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/patilsuraj767/turf/turf/model"
)

func GetCustomer(c echo.Context) error {
	var customer model.Customer
	err := customer.GetCustomerById(c.Param("id"))
	if err != nil {
		return c.JSON(500, echo.Map{"status": "error", "message": err})
	} else {
		return c.JSON(200, customer)
	}
}

func GetCustomers(c echo.Context) error {
	var customer model.Customer
	customers, err := customer.GetAllCustomers()
	if err != nil {
		return c.JSON(500, echo.Map{"status": "error", "message": err})
	} else {
		return c.JSON(200, customers)
	}
}
