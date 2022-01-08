package model

import (
	"errors"

	"github.com/patilsuraj767/turf/turf/db"
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	Name   string
	Mobile string `gorm:"unique"`
}

func (b *Customer) GetCustomerById(id string) error {
	err := db.DBConn.First(&b, id)
	if errors.Is(err.Error, gorm.ErrRecordNotFound) {
		return errors.New("NO RECORD FOUND")
	} else {
		return nil
	}
}

func (Customer) GetAllCustomers() ([]Customer, error) {
	var customerList []Customer
	err := db.DBConn.Find(&customerList)
	if errors.Is(err.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("NO RECORD FOUND")
	} else {
		return customerList, nil
	}
}

func (c *Customer) GetCustomerByMobile(mobile string) *gorm.DB {
	err := db.DBConn.Where("Mobile = ?", mobile).First(&c)
	return err
}

func (c *Customer) CreateCustomer() *gorm.DB {
	result := db.DBConn.Create(&c)
	return result
}
