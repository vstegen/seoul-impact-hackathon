package food

type RestaurantStatus string

const (
	StatusOpen   RestaurantStatus = "Open"
	StatusClosed RestaurantStatus = "Closed"
)

type FoodOptions string

const (
	FoodOptionVegan      FoodOptions = "Vegan"
	FoodOptionVegetarian FoodOptions = "Vegetarian"
	FoodOptionHalal      FoodOptions = "Halal"
	FoodOptionGlutenFree FoodOptions = "Gluten Free"
)

type Facility struct {
	Id            int              `json:"id"`
	OpeningTime   string           `json:"openingTime"`
	Rating        *int             `json:"rating,omitempty"`
	Desc          *string          `json:"desc,omitempty"`
	Address       string           `json:"address"`
	Name          string           `json:"name"`
	CurrentStatus RestaurantStatus `json:"currentStatus"`
	Announcement  *string          `json:"announcement,omitempty"`
	Contact       *string          `json:"contact,omitempty"`
	Website       *string          `json:"website,omitempty"`
	FoodTypes     []FoodOptions    `json:"foodTypes"`
}
