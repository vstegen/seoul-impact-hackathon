package food

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type RestaurantStatus string

const (
	StatusOpen   RestaurantStatus = "Open"
	StatusClosed RestaurantStatus = "Closed"
)

type Facility struct {
	Id            int              `json: "id"`
	OpeningTime   string           `json: "openingTime"`
	Rating        int              `json: "rating", omitempty`
	Desc          string           `json: "desc", omitempty`
	Address       string           `json: "address"`
	Name          string           `json: "name"`
	CurrentStatus RestaurantStatus `json: "currentStatus"`
	Announcement  string           `json: "announcement", omitempty`
	Contact       string           `json: "contact", omitempty`
	Website       string           `json: "website", omitempty`
}

// TODO: Use a db instead of hardcoding
var facilities = []Facility{
	{
		Id:            1,
		Name:          "Korea Women’s Hotline",
		Address:       "8-4, Jinheung-ro 16-gil, Eunpyeong-gu, Seoul, 03369, Republic of KOREA",
		Contact:       "82-2-3156-5400",
		Website:       "enghotline.cafe24.com",
		OpeningTime:   "Mon-Fri 20:00-10:00, Sat-Sun 18:00-10:00, Holidays 18:00-10:00",
		CurrentStatus: StatusOpen,
	},
	{
		Id:            2,
		Name:          "Salvation Army Korea",
		Address:       "130 Deoksugung-gil, Jung-gu, Seoul 100-120",
		Contact:       "+82-2-6364-4000",
		Website:       "en.salvationarmy.kr",
		OpeningTime:   "Mon-Fri 20:00-10:00, Sat-Sun 18:00-10:00, Holidays 18:00-10:00",
		CurrentStatus: StatusClosed,
	},
}

func GetFacilities(c echo.Context) error {
	return c.JSON(http.StatusOK, facilities)
}

func GetFacility(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Errorf("invalid id '%s': %w", idParam, err))
	}
	// TODO: use a map instead of looping if significantly more facilities are added
	for _, fac := range facilities {
		if id == fac.Id {
			return c.JSON(http.StatusOK, fac)
		}
	}

	return echo.NewHTTPError(http.StatusNotFound, nil)
}
