package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"youth-korea/food"
	"youth-korea/shelters"
)

func main() {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAccessControlAllowOrigin},
	}))

	shelterRoutes := shelters.NewRoutes(shelters.ShelterRepo{})
	foodRoutes := food.NewRoutes(food.FacilityRepo{})

	e.GET("/shelters", shelterRoutes.GetShelters)
	e.GET("/shelters/:id", shelterRoutes.GetShelter)
	e.GET("/food", foodRoutes.GetFacilities)
	e.GET("/food/:id", foodRoutes.GetFacility)

	e.Logger.Fatal(e.Start(":8000"))
}
