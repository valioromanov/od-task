package internal

type Price struct {
	Day int `json:"day"`
}

type VehicleInfo struct {
	VehicleID       int    `json:"id"`
	Name            string `json:"name"`
	Description     string `json:"description"`
	Type            string `json:"type"`
	Make            string `json:"make"`
	Model           string `json:"model"`
	Year            string `json:"year"`
	Length          string `json:"length"`
	Sleeps          string `json:"sleeps"`
	PrimaryImageUrl string `json:"primary_image_url"`
	Price           Price  `json:"price"`
}

type Location struct {
	City       string  `json:"city"`
	State      string  `json:"state"`
	Zip        string  `json:"zip"`
	Country    string  `json:"country"`
	Latitude   float64 `json:"lat"`
	Longtitude float64 `json:"lng"`
}

type User struct {
	UserID    int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type GetRentalResponse struct {
	VehicleInfo
	Location Location `json:"location"`
	User     User     `json:"user"`
}
