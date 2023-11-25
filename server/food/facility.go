package food

type RestaurantStatus string

const (
	StatusOpen   RestaurantStatus = "Open"
	StatusClosed RestaurantStatus = "Closed"
)

type FoodOption string

const (
	FoodOptionVegan      FoodOption = "Vegan"
	FoodOptionVegetarian FoodOption = "Vegetarian"
	FoodOptionHalal      FoodOption = "Halal"
	FoodOptionGlutenFree FoodOption = "Gluten Free"
)

// NOTE: pointer fields are used for clear separation between
// required and optional fields. Any field that is not a pointer
// is expected to be available information about the facility.
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
	FoodTypes     []FoodOption     `json:"foodTypes"`
}
