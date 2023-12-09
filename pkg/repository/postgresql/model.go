package postgresql

import (
	"time"
)

type Rentals struct {
	ID              int       `gorm:"primarykey column:id"`
	UserForeignID   string    `gorm:"column:user_id"`
	Name            string    `gorm:"column:name"`
	Type            string    `gorm:"column:type"`
	Description     string    `gorm:"column:description"`
	Sleeps          int       `gorm:"column:sleep"`
	PricePerDay     int64     `gorm:"column:price_per_day"`
	HomeCity        string    `gorm:"column:home_city"`
	HomeState       string    `gorm:"column:home_state"`
	HomeZip         string    `gorm:"column:home_zip"`
	HomeCountry     string    `gorm:"column:home_country"`
	VehicleMake     string    `gorm:"column:vehicle_make"`
	VehicleModel    string    `gorm:"column:vehicle_model"`
	VehicleYear     int       `gorm:"column:vehicle_year"`
	VehicleLength   float64   `gorm:"column:vehicle_length"`
	Created         time.Time `gorm:"column:created"`
	Updated         time.Time `gorm:"column:updated"`
	Latitude        float64   `gorm:"column:lat"`
	Longtitude      float64   `gorm:"column:lng"`
	PrimaryImageUrl string    `gorm:"column:primary_image_url"`
}

type Users struct {
	UserID    int    `gorm:"primarykey; column:id"`
	FirstName string `gorm:"column:first_name"`
	LastName  string `gorm:"column:last_name"`
}

type FindResult struct {
	Rentals
	Users
}
