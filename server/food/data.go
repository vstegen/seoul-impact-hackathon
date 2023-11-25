package food

import (
	"errors"

	"youth-korea/utils"
)

var errNotFound = errors.New("shelter not found")

// TODO: Use a db instead of hardcoding
var facilities = []Facility{
	{
		Id:            1,
		Name:          "Seoul Seafood",
		Address:       "8-4, Jinheung-ro 16-gil, Eunpyeong-gu, Seoul, 03369, Republic of KOREA",
		Contact:       utils.String("82-2-3156-5400"),
		Website:       utils.String("enghotline.cafe24.com"),
		OpeningTime:   "Mon-Fri 20:00-10:00, Sat-Sun 18:00-10:00, Holidays 18:00-10:00",
		CurrentStatus: StatusOpen,
	},
	{
		Id:            2,
		Name:          "Gukbap Alley",
		Address:       "130 Deoksugung-gil, Jung-gu, Seoul 100-120",
		Contact:       utils.String("+82-2-6364-4000"),
		Website:       utils.String("en.salvationarmy.kr"),
		OpeningTime:   "Mon-Fri 20:00-10:00, Sat-Sun 18:00-10:00, Holidays 18:00-10:00",
		CurrentStatus: StatusClosed,
	},
	{
		Id:            3,
		Name:          "7-Eleven",
		Address:       "8-4, Jinheung-ro 16-gil, Eunpyeong-gu, Seoul, 03369, Republic of KOREA",
		Contact:       utils.String("82-2-3156-5400"),
		Website:       utils.String("enghotline.cafe24.com"),
		OpeningTime:   "Mon-Fri 20:00-10:00, Sat-Sun 18:00-10:00, Holidays 18:00-10:00",
		CurrentStatus: StatusOpen,
		FoodTypes:     []FoodOption{FoodOptionVegan, FoodOptionVegetarian},
	},
	{
		Id:            4,
		Name:          "Salvation Army Korea",
		Address:       "130 Deoksugung-gil, Jung-gu, Seoul 100-120",
		Contact:       utils.String("+82-2-6364-4000"),
		Website:       utils.String("en.salvationarmy.kr"),
		OpeningTime:   "Mon-Fri 20:00-10:00, Sat-Sun 18:00-10:00, Holidays 18:00-10:00",
		CurrentStatus: StatusClosed,
		FoodTypes:     []FoodOption{FoodOptionHalal, FoodOptionVegetarian},
	},
}

type FacilityRepo struct{}

// NOTE: It is intended for this method to always return a nil error.
// This is done for extensibility later when the underlying data source
// changes into a potentially failing one.
func (r FacilityRepo) Get() ([]Facility, error) {
	return facilities, nil
}

func (r FacilityRepo) GetById(id int) (*Facility, error) {
	allFacilities, err := r.Get()
	if err != nil {
		return nil, err
	}

	for _, f := range allFacilities {
		if f.Id == id {
			return &f, nil
		}
	}

	return nil, errNotFound
}

func (r FacilityRepo) GetBy(f filter) ([]Facility, error) {
	allFacilities, err := r.Get()
	if err != nil {
		return nil, errNotFound
	}

	return f.apply(allFacilities), nil
}
