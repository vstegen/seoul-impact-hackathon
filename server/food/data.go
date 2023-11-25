package food

import "youth-korea/utils"

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
