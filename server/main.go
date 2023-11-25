package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/shelters", getShelters)
	e.Logger.Fatal(e.Start(":8000"))
}

type ShelterStatus string

const (
	StatusOpen   ShelterStatus = "Open"
	StatusClosed ShelterStatus = "Closed"
	StatusFull   ShelterStatus = "Full"
)

type Shelter struct {
	OpeningTime   string        `json: "openingTime"`
	Facilities    []string      `json: "facilities", omitempty`
	Requirements  string        `json: "requirements", omitempty`
	Rules         string        `json: "rules"`
	MaxCapacity   int           `json: "maxCapacity"`
	Capacity      int           `json: "capacity"`
	Rating        int           `json: "rating", omitempty`
	Desc          string        `json: "desc", omitempty`
	Address       string        `json: "address"`
	Name          string        `json: "name"`
	CurrentStatus ShelterStatus `json: "currentStatus"`
	Announcement  string        `json: "announcement", omitempty`
	Contact       string        `json: "contact", omitempty`
	Website       string        `json: "website", omitempty`
}

var shelters = []Shelter{
	{
		Name:          "Korea Women’s Hotline",
		Address:       "8-4, Jinheung-ro 16-gil, Eunpyeong-gu, Seoul, 03369, Republic of KOREA",
		Contact:       "82-2-3156-5400",
		Website:       "enghotline.cafe24.com",
		Requirements:  "Under 18 years old",
		OpeningTime:   "Mon-Fri 20:00-10:00, Sat-Sun 18:00-10:00, Holidays 18:00-10:00",
		Rules:         "<Imagine some rules here>",
		MaxCapacity:   20,
		Capacity:      5,
		CurrentStatus: StatusOpen,
	},
	{
		Name:          "Salvation Army Korea",
		Address:       "130 Deoksugung-gil, Jung-gu, Seoul 100-120",
		Facilities:    []string{"Childcare", "Physical Health", "Education", "Rent Assistance", "Utility Assistance", "Meals", "Emergency Shelters", "Family Shelters", "Youth Shelters", "Counseling", "Clothing"},
		Contact:       "+82-2-6364-4000",
		Website:       "en.salvationarmy.kr",
		OpeningTime:   "Mon-Fri 20:00-10:00, Sat-Sun 18:00-10:00, Holidays 18:00-10:00",
		Rules:         "<Imagine some rules here>",
		MaxCapacity:   15,
		Capacity:      0,
		CurrentStatus: StatusClosed,
	},
	{
		Name:          "Anna’s House",
		Address:       "118 Hadaewon-dong, Jungian-gu Seongnam-si, Gyeonggi-do",
		Facilities:    []string{"Emergency Shelters", "Meals", "Clothing", "Physical Health", "Counseling", "Education"},
		Contact:       "031-757-6336",
		Website:       "annahouse.or.kr",
		OpeningTime:   "Mon-Fri 20:00-10:00, Sat-Sun 18:00-10:00, Holidays 18:00-10:00",
		Rules:         "<Imagine some rules here>",
		MaxCapacity:   6,
		Capacity:      6,
		CurrentStatus: StatusFull,
		Announcement:  "We are currently full, please check back later",
	},
	{
		Name:          "SOS Children's Village Seoul",
		Address:       "Seoul, South Korea",
		Facilities:    []string{"Childcare", "Youth Shelters", "Meals", "Physical Health", "Education"},
		Contact:       "02-3453-8400",
		Website:       "sos-childrensvillages.org",
		OpeningTime:   "Mon-Fri 20:00-10:00, Sat-Sun 18:00-10:00, Holidays 18:00-10:00",
		Rules:         "<Imagine some rules here>",
		MaxCapacity:   20,
		Capacity:      5,
		CurrentStatus: StatusOpen,
	},
	{
		Name:          "Happy Home",
		Address:       "675-2 Bupyeong 2(i)-dong, Bupyeong-gu, Incheon, South Korea",
		Facilities:    []string{"Childcare", "Youth Shelters", "Meals", "Education", "Counseling"},
		Contact:       "+82 32-518-2080",
		Website:       "happyhome.or.kr",
		OpeningTime:   "Mon-Fri 20:00-10:00, Sat-Sun 18:00-10:00, Holidays 18:00-10:00",
		Rules:         "<Imagine some rules here>",
		MaxCapacity:   20,
		Capacity:      5,
		CurrentStatus: StatusOpen,
	},
	{
		Name:          "Rainbow Youth Center",
		Address:       "20 Jahamun-ro 24-gil, Jongno-gu, Seoul, South Korea",
		Facilities:    []string{"Youth Shelters", "Counseling", "Education"},
		Contact:       "02-733-7587",
		Website:       "rainbowyouth.or.kr",
		OpeningTime:   "Mon-Fri 20:00-10:00, Sat-Sun 18:00-10:00, Holidays 18:00-10:00",
		Rules:         "<Imagine some rules here>",
		MaxCapacity:   20,
		Capacity:      5,
		CurrentStatus: StatusOpen,
	},
}

func getShelters(c echo.Context) error {
	return c.JSON(http.StatusOK, shelters)
}
