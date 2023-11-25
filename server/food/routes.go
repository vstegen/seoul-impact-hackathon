package food

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

// TODO: potentially add query param validation for the filter
func GetFacilities(c echo.Context) error {
	activeFilter := filter{}
	activeFilter.status = RestaurantStatus(c.QueryParam("status"))

	foodOptionsQueryParam := c.QueryParam("foodOptions")
	if foodOptionsQueryParam != "" {
		foodOptions := strings.Split(foodOptionsQueryParam, ",")
		activeFilter.foodOptions = make(map[string]struct{})
		for _, foodOption := range foodOptions {
			activeFilter.foodOptions[strings.ToLower(foodOption)] = struct{}{}
		}
	}

	filteredFacilities := activeFilter.apply(facilities)
	if len(filteredFacilities) == 0 {
		return c.JSON(http.StatusOK, []Facility{})
	}

	return c.JSON(http.StatusOK, filteredFacilities)
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
