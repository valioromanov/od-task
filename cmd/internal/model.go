package internal

type Price struct {
	Day int64 `json:"day"`
}

type VehicleInfo struct {
	VehicleID       uint    `json:"id"`
	Name            string  `json:"name"`
	Description     string  `json:"description"`
	Type            string  `json:"type"`
	Make            string  `json:"make"`
	Model           string  `json:"model"`
	Year            int     `json:"year"`
	Length          float64 `json:"length"`
	Sleeps          int     `json:"sleeps"`
	PrimaryImageUrl string  `json:"primary_image_url"`
	Price           Price   `json:"price"`
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
