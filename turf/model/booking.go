package model

import (
	"errors"
	"time"

	"github.com/patilsuraj767/turf/turf/db"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Booking struct {
	gorm.Model
	CustomerID  uint
	Customer    Customer
	BookingDate datatypes.Date `gorm:"not null"`
	StartTime   datatypes.Time `gorm:"not null"`
	EndTime     datatypes.Time `gorm:"not null"`
	Description string
}

func (b *Booking) DeleteBookingById() error {
	err := db.DBConn.Delete(&b)
	if errors.Is(err.Error, gorm.ErrRecordNotFound) {
		return errors.New("NO RECORD FOUND")
	} else {
		return nil
	}
}

func (b *Booking) GetBookingsById(id string) error {
	err := db.DBConn.Preload("Customer").First(&b, id)
	if errors.Is(err.Error, gorm.ErrRecordNotFound) {
		return errors.New("NO RECORD FOUND")
	} else {
		return nil
	}
}

func (Booking) GetBookingsByDate(t time.Time) ([]Booking, error) {
	var bookingList []Booking
	inputDate := datatypes.Date(t)
	err := db.DBConn.Where("booking_date = ?", inputDate).Preload("Customer").Order("start_time").Find(&bookingList)
	if errors.Is(err.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("NO RECORD FOUND")
	} else {
		return bookingList, nil
	}
}

func (b *Booking) CreateBooking() (*gorm.DB, error) {
	var booking Booking
	result := db.DBConn.Where(
		"start_time < ? AND end_time > ?", b.StartTime, b.StartTime).Or(
		"start_time < ? AND end_time > ?", b.EndTime, b.EndTime).Or(
		"start_time > ? AND end_time < ?", b.StartTime, b.EndTime).Or(
		"start_time > ? AND start_time < ?", b.StartTime, b.EndTime).First(&booking)
	if result.Error == gorm.ErrRecordNotFound {
		result := db.DBConn.Create(&b)
		if result.Error == nil {
			return result, nil
		} else {
			return nil, errors.New("SOMETHING WENT WRONG, UNABLE TO CREATE BOOKING")
		}
	} else {
		return nil, errors.New("Booking already persent for given slot")
	}
}

func (b *Booking) EditBooking() (*gorm.DB, error) {
	var booking Booking
	result := db.DBConn.Where(
		"start_time < ? AND end_time > ?", b.StartTime, b.StartTime).Or(
		"start_time < ? AND end_time > ?", b.EndTime, b.EndTime).Or(
		"start_time > ? AND end_time < ?", b.StartTime, b.EndTime).Or(
		"start_time > ? AND start_time < ?", b.StartTime, b.EndTime).First(&booking)
	if result.Error == gorm.ErrRecordNotFound {
		result := db.DBConn.Save(&b)
		if result.Error == nil {
			return result, nil
		} else {
			return nil, errors.New("SOMETHING WENT WRONG, UNABLE TO CREATE BOOKING")
		}
	} else {
		return nil, errors.New("Booking already persent for given slot")
	}
}
