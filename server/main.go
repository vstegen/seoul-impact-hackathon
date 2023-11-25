package main

import (
	"github.com/labstack/echo/v4"

	"youth-korea/food"
	"youth-korea/shelters"
)

func main() {
	e := echo.New()
	e.GET("/shelters", shelters.GetShelters)
	e.GET("/shelters/:id", shelters.GetShelter)
	e.GET("/food", food.GetFacilities)
	e.GET("/food/:id", food.GetFacility)
	e.Logger.Fatal(e.Start(":8000"))
}
