package internal

type Price struct {
	Day int64 `json:"day"`
}

type VehicleInfo struct {
	VehicleID       int     `json:"id"`
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

type Filters struct {
	MinPrice int64  `form:"price_min"`
	MaxPrice int64  `form:"price_max"`
	Limit    uint   `form:"limit"`
	Offset   uint   `form:"offset,default=0"`
	IDs      string `form:"ids"`
	Near     string `form:"near"` //[lat,lng]
	Sort     string `form:"sort"`
}
