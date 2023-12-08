package postgresql

import "time"

type Rentals struct {
	RentalID        string    `gorm:"primarykey"`
	UserID          string    `gorm:"user_id"`
	Name            string    `gorm:"name"`
	Type            string    `gorm:"type"`
	Description     string    `gorm:"description"`
	Sleeps          int       `gorm:"sleep"`
	PricePerDay     int64     `gorm:"price_per_day"`
	HomeCity        string    `gorm:"home_city"`
	HomeState       string    `gorm:"home_state"`
	HomeZip         string    `gorm:"home_zip"`
	HomeCountry     string    `gorm:"home_country"`
	VehicleMake     string    `gorm:"vehicle_make"`
	VehicleMode     string    `gorm:"vehicle_model"`
	VehicleYear     int       `gorm:"vehicle_year"`
	VehicleLength   float64   `gorm:"vehicle_length"`
	Created         time.Time `gorm:"created"`
	Updated         time.Time `gorm:"updated"`
	Lat             float32   `gorm:"lat"`
	Lng             float32   `gorm:"lng"`
	PrimaryImageUrl string    `gorm:"primary_image_url"`
}

type Users struct {
	UserID    string `gorm:"primarykey"`
	FirstName string `gorm:"first_name"`
	LastName  string `gorm:"last_name"`
}

type FindResult struct {
	Rentals
	Users
}
